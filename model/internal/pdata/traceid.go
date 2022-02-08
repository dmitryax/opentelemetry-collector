// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pdata // import "go.opentelemetry.io/collector/model/pdata"

import (
	"go.opentelemetry.io/collector/model/internal/data"
	otlplogs "go.opentelemetry.io/collector/model/internal/data/protogen/logs/v1"
	otlpmetrics "go.opentelemetry.io/collector/model/internal/data/protogen/metrics/v1"
)

// TraceID is an alias of OTLP TraceID data type.
type TraceID struct {
	orig data.TraceID
}

// InvalidTraceID returns an empty (all zero bytes) TraceID.
func InvalidTraceID() TraceID {
	return TraceID{orig: data.NewTraceID([16]byte{})}
}

// TraceIDFromBytes returns a new TraceID from the given byte array.
func TraceIDFromBytes(bytes [16]byte) TraceID {
	return TraceID{orig: data.NewTraceID(bytes)}
}

func NewTraceID(orig data.TraceID) TraceID {
	return TraceID{orig: orig}
}

// Bytes returns the byte array representation of the TraceID.
func (t TraceID) Bytes() [16]byte {
	return t.orig.Bytes()
}

// HexString returns hex representation of the TraceID.
func (t TraceID) HexString() string {
	return t.orig.HexString()
}

// IsEmpty returns true if id doesn't contain at least one non-zero byte.
func (t TraceID) IsEmpty() bool {
	return t.orig.IsEmpty()
}

func (t TraceID) CopyToLogRecord(origLog *otlplogs.LogRecord) {
	origLog.TraceId = t.orig
}

func (t TraceID) CopyToExemplar(origLog *otlpmetrics.Exemplar) {
	origLog.TraceId = t.orig
}
