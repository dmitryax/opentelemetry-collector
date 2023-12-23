// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exporterqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueConfig_Validate(t *testing.T) {
	qCfg := NewDefaultConfig()
	assert.NoError(t, qCfg.Validate())

	qCfg.NumConsumers = 0
	assert.EqualError(t, qCfg.Validate(), "number of consumers must be positive")

	qCfg = NewDefaultConfig()
	qCfg.QueueSizeItems = 0
	assert.EqualError(t, qCfg.Validate(), "queue size must be positive")

	qCfg.QueueSize = 1
	qCfg.QueueSizeItems = 1
	assert.EqualError(t, qCfg.Validate(), "only one of 'queue_size' and 'queue_size_items' can be specified")

	// Confirm Validate doesn't return error with invalid config when feature is disabled
	qCfg.Enabled = false
	assert.NoError(t, qCfg.Validate())
}
