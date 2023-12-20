// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exporterhelper // import "go.opentelemetry.io/collector/exporter/exporterhelper"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal"
)

// QueueConfig defines configuration for queueing requests before exporting.
// It's supposed to be used with the new exporter helpers New[Traces|Metrics|Logs]RequestExporter.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
type QueueConfig struct {
	// Enabled indicates whether to not enqueue batches before exporting.
	Enabled bool `mapstructure:"enabled"`
	// NumConsumers is the number of consumers from the queue.
	NumConsumers int `mapstructure:"num_consumers"`
	// QueueItemsSize is the maximum number of items (spans, metric data points or log records)
	// allowed in queue at any given time.
	QueueItemsSize int `mapstructure:"queue_size"`
}

// NewDefaultQueueConfig returns the default Config.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
func NewDefaultQueueConfig() QueueConfig {
	return QueueConfig{
		Enabled:        true,
		NumConsumers:   10,
		QueueItemsSize: 100_000,
	}
}

// Validate checks if the QueueSettings configuration is valid
func (qCfg *QueueConfig) Validate() error {
	if !qCfg.Enabled {
		return nil
	}
	if qCfg.NumConsumers <= 0 {
		return errors.New("number of consumers must be positive")
	}
	if qCfg.QueueItemsSize <= 0 {
		return errors.New("queue size must be positive")
	}
	return nil
}

// Queue defines a producer-consumer exchange which can be backed by e.g. the memory-based ring buffer queue
// (boundedMemoryQueue) or via a disk-based queue (persistentQueue)
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
type Queue = internal.Queue[Request]

// QueueCapacityUnit defines the type of capacity limiter for the queue.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
type QueueCapacityUnit string

const (
	// QueueCapacityUnitRequest indicates that the queue capacity is measured in requests.
	QueueCapacityUnitRequest QueueCapacityUnit = "request"
	// QueueCapacityUnitItem indicates that the queue capacity is measured in items (spans,
	// metric data points or log records).
	QueueCapacityUnitItem QueueCapacityUnit = "item"
)

type QueueCapacityLimiter = internal.QueueCapacityLimiter[Request]

type QueueCreateSettings struct {
	CreateSettings  exporter.CreateSettings
	CapacityLimiter QueueCapacityLimiter
	DataType        component.DataType
}

type QueueFactory func(QueueCreateSettings) Queue

// NewMemoryQueueFactory returns a factory to create a new memory queue.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
func NewMemoryQueueFactory() QueueFactory {
	return func(set QueueCreateSettings) Queue {
		return internal.NewBoundedMemoryQueue[Request](set.CapacityLimiter)
	}
}

// PersistentQueueConfig defines configuration for queueing requests in a persistent storage.
// The struct is provided to be added in the exporter configuration as one struct under the "sending_queue" key.
// The exporter helper Go interface requires the fields to be provided separately to WithRequestQueue and
// NewPersistentQueueFactory.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
type PersistentQueueConfig struct {
	QueueConfig `mapstructure:",squash"`
	// StorageID if not empty, enables the persistent storage and uses the component specified
	// as a storage extension for the persistent queue
	StorageID *component.ID `mapstructure:"storage"`
}

// NewPersistentQueueFactory returns a factory to create a new persistent queue.
// If cfg.StorageID is nil then it falls back to memory queue.
// This API is at the early stage of development and may change without backward compatibility
// until https://github.com/open-telemetry/opentelemetry-collector/issues/8122 is resolved.
func NewPersistentQueueFactory(storageID *component.ID, marshaler RequestMarshaler, unmarshaler RequestUnmarshaler) QueueFactory {
	if storageID == nil {
		return NewMemoryQueueFactory()
	}
	return func(set QueueCreateSettings) Queue {
		return internal.NewPersistentQueue[Request](set.CapacityLimiter, set.DataType, *storageID, marshaler, unmarshaler, set.CreateSettings)
	}
}
