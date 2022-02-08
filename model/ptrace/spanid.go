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

package ptrace // import "go.opentelemetry.io/collector/model/pdata"

import (
	"go.opentelemetry.io/collector/model/internal/pdata"
)

// SpanID is an alias of OTLP SpanID data type.
type SpanID struct {
	pdata.SpanID
}

// InvalidSpanID returns an empty (all zero bytes) SpanID.
func InvalidSpanID() SpanID {
	return SpanID{pdata.InvalidSpanID()}
}

// NewSpanID returns a new SpanID from the given byte array.
func NewSpanID(bytes [8]byte) SpanID {
	return SpanID{pdata.SpanIDFromBytes(bytes)}
}
