// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package queuebatch // import "go.opentelemetry.io/collector/exporter/exporterhelper/internal/queuebatch"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"
)

type sharder struct {
	bCfg        BatchConfig
	bSet        batcherSettings[request.Request]
	partitioner Partitioner[request.Request]
	workerPool  chan struct{} // Worker pool for limiting concurrency, nil if no limit is set.
	singleShard *shardBatcher // For single shard case, we use a single shard batcher.
	shards      sync.Map      // For multi shard case, we use a map of shard batchers.
}

func newSharder(bCfg BatchConfig, bSet batcherSettings[request.Request], partitioner Partitioner[request.Request], maxWorkers int) *sharder {
	s := &sharder{
		bCfg:        bCfg,
		bSet:        bSet,
		partitioner: partitioner,
	}
	if maxWorkers != 0 {
		s.workerPool = make(chan struct{}, maxWorkers)
		for i := 0; i < maxWorkers; i++ {
			s.workerPool <- struct{}{}
		}
	}
	if partitioner == nil {
		s.singleShard = newShardBatcher(bCfg, bSet, s.workerPool)
	}
	return s
}

func (sh *sharder) getShard(ctx context.Context, req request.Request) *shardBatcher {
	if sh.singleShard != nil {
		// If we have a single shard, return it directly.
		return sh.singleShard
	}

	key := sh.partitioner.GetKey(ctx, req)
	if shard, ok := sh.shards.Load(key); ok {
		return shard.(*shardBatcher)
	}
	sb := newShardBatcher(sh.bCfg, sh.bSet, sh.workerPool)
	sb.start()
	sh.shards.Store(key, sb)
	return sb
}

func (sh *sharder) Start(context.Context, component.Host) error {
	if sh.singleShard != nil {
		sh.singleShard.start()
	}
	return nil
}

func (sh *sharder) Shutdown(context.Context) error {
	if sh.singleShard != nil {
		sh.singleShard.shutdown()
	}

	stopWG := sync.WaitGroup{}
	sh.shards.Range(func(key, shard any) bool {
		stopWG.Add(1)
		go func() {
			shard.(*shardBatcher).shutdown()
			stopWG.Done()
		}()
		return true
	})
	stopWG.Wait()
	return nil
}
