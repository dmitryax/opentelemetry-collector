// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exporterqueue // import "go.opentelemetry.io/collector/exporter/exporterqueue"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/internal"
)

// Queue defines a producer-consumer exchange which can be backed by e.g. the memory-based ring buffer queue
// (boundedMemoryQueue) or via a disk-based queue (persistentQueue)
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
type Queue[T any] internal.Queue[T]

// Settings defines settings for creating a queue.
type Settings[T any] struct {
	Sizer            internal.Sizer[T]
	Capacity         int
	DataType         component.DataType
	ExporterSettings exporter.CreateSettings
}

type Factory[T any] func(Settings[T]) Queue[T]

// NewMemoryQueueFactory returns a factory to create a new memory queue.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
func NewMemoryQueueFactory[T any]() Factory[T] {
	return func(set Settings[T]) Queue[T] {
		return internal.NewBoundedMemoryQueue[T](internal.MemoryQueueSettings[T]{
			Sizer:    set.Sizer,
			Capacity: set.Capacity,
		})
	}
}

// PersistentQueueSettings defines developer settings for the persistent queue factory.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
type PersistentQueueSettings[T any] struct {
	// Marshaler is used to serialize queue elements before storing them in the persistent storage.
	Marshaler func(req T) ([]byte, error)
	// Unmarshaler is used to deserialize requests after reading them from the persistent storage.
	Unmarshaler func(data []byte) (T, error)
}

// NewPersistentQueueFactory returns a factory to create a new persistent queue.
// If cfg.StorageID is nil then it falls back to memory queue.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
func NewPersistentQueueFactory[T any](storageID *component.ID, factorySettings PersistentQueueSettings[T]) Factory[T] {
	if storageID == nil {
		return NewMemoryQueueFactory[T]()
	}
	return func(set Settings[T]) Queue[T] {
		return internal.NewPersistentQueue[T](internal.PersistentQueueSettings[T]{
			Sizer:            set.Sizer,
			Capacity:         set.Capacity,
			DataType:         set.DataType,
			StorageID:        *storageID,
			Marshaler:        factorySettings.Marshaler,
			Unmarshaler:      factorySettings.Unmarshaler,
			ExporterSettings: set.ExporterSettings,
		})
	}
}
