// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/exporter/exporterhelper/internal"

import (
	"sync/atomic"
)

type itemsCounter interface {
	ItemsCount() int
}

// QueueCapacityLimiter is an interface to control the capacity of the queue.
type QueueCapacityLimiter[T any] interface {
	// Capacity is the maximum capacity of the queue.
	Capacity() int
	// Size returns the current size of the queue.
	Size() int
	// Claim tries to claim capacity for the given element. If the capacity is not available, it returns false.
	Claim(T) bool
	// Release releases capacity for the given queue element.
	Release(T)
	// SizeOf returns the size of the given element.
	SizeOf(T) uint64
}

type baseCapacityLimiter[T any] struct {
	used *atomic.Uint64
	cap  uint64
}

func newBaseCapacityLimiter[T any](capacity int) baseCapacityLimiter[T] {
	return baseCapacityLimiter[T]{
		used: &atomic.Uint64{},
		cap:  uint64(capacity),
	}
}

func (bcl *baseCapacityLimiter[T]) Capacity() int {
	return int(bcl.cap)
}

func (bcl *baseCapacityLimiter[T]) Size() int {
	return int(bcl.used.Load())
}

//nolint:unused
func (bcl *baseCapacityLimiter[T]) claim(claimed uint64) bool {
	if bcl.used.Load()+claimed > bcl.cap {
		return false
	}
	bcl.used.Add(claimed)
	return true
}

//nolint:unused
func (bcl *baseCapacityLimiter[T]) release(capacity uint64) {
	bcl.used.Add(^(capacity - 1))
}

// itemsCapacityLimiter is a capacity limiter that limits the queue based on the number of items (e.g. spans, log records).
type itemsCapacityLimiter[T itemsCounter] struct {
	baseCapacityLimiter[T]
}

func NewItemsCapacityLimiter[T itemsCounter](capacity int) QueueCapacityLimiter[T] {
	return &itemsCapacityLimiter[T]{baseCapacityLimiter: newBaseCapacityLimiter[T](capacity)}
}

// nolint:unused
func (icl *itemsCapacityLimiter[T]) Claim(el T) bool {
	return icl.baseCapacityLimiter.claim(uint64(el.ItemsCount()))
}

// nolint:unused
func (icl *itemsCapacityLimiter[T]) Release(el T) {
	icl.baseCapacityLimiter.release(uint64(el.ItemsCount()))
}

// nolint:unused
func (icl *itemsCapacityLimiter[T]) SizeOf(el T) uint64 {
	return uint64(el.ItemsCount())
}

var _ QueueCapacityLimiter[itemsCounter] = (*itemsCapacityLimiter[itemsCounter])(nil)

// requestsCapacityLimiter is a capacity limiter that limits the queue based on the number of requests.
type requestsCapacityLimiter[T any] struct {
	baseCapacityLimiter[T]
}

func NewRequestsCapacityLimiter[T any](capacity int) QueueCapacityLimiter[T] {
	return &requestsCapacityLimiter[T]{baseCapacityLimiter: newBaseCapacityLimiter[T](capacity)}
}

// nolint:unused
func (rcl *requestsCapacityLimiter[T]) Claim(_ T) bool {
	return rcl.baseCapacityLimiter.claim(1)
}

// nolint:unused
func (rcl *requestsCapacityLimiter[T]) Release(_ T) {
	rcl.baseCapacityLimiter.release(1)
}

// nolint:unused
func (rcl *requestsCapacityLimiter[T]) SizeOf(_ T) uint64 {
	return 1
}

var _ QueueCapacityLimiter[any] = (*requestsCapacityLimiter[any])(nil)
