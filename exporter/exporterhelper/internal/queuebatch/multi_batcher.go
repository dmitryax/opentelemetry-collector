// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package queuebatch // import "go.opentelemetry.io/collector/exporter/exporterhelper/internal/queuebatch"

import (
	"context"

	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"
)

type multiBatcher struct {
	*sharder
}

var _ Batcher[request.Request] = (*multiBatcher)(nil)

func newMultiBatcher(sharder *sharder) multiBatcher {
	return multiBatcher{sharder: sharder}
}

func (qb multiBatcher) Consume(ctx context.Context, req request.Request, done Done) {
	qb.sharder.getShard(ctx, req).Consume(ctx, req, done)
}
