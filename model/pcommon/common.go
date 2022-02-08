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

package pcommon // import "go.opentelemetry.io/collector/model/pdata"

// This file contains data structures that are common for all telemetry types,
// such as timestamps, attributes, etc.

import (
	otlpcommon "go.opentelemetry.io/collector/model/internal/data/protogen/common/v1"
	"go.opentelemetry.io/collector/model/internal/pdata"
)

// AttributeValueType specifies the type of AttributeValue.
type AttributeValueType pdata.AttributeValueType

const (
	AttributeValueTypeEmpty  = AttributeValueType(pdata.AttributeValueTypeEmpty)
	AttributeValueTypeString = AttributeValueType(pdata.AttributeValueTypeString)
	AttributeValueTypeInt    = AttributeValueType(pdata.AttributeValueTypeInt)
	AttributeValueTypeDouble = AttributeValueType(pdata.AttributeValueTypeDouble)
	AttributeValueTypeBool   = AttributeValueType(pdata.AttributeValueTypeBool)
	AttributeValueTypeMap    = AttributeValueType(pdata.AttributeValueTypeMap)
	AttributeValueTypeArray  = AttributeValueType(pdata.AttributeValueTypeArray)
	AttributeValueTypeBytes  = AttributeValueType(pdata.AttributeValueTypeBytes)
)

// String returns the string representation of the AttributeValueType.
func (avt AttributeValueType) String() string {
	return pdata.AttributeValueType(avt).String()
}

// AttributeValue is a mutable cell containing the value of an attribute. Typically used in AttributeMap.
// Must use one of NewAttributeValue+ functions below to create new instances.
//
// Intended to be passed by value since internally it is just a pointer to actual
// value representation. For the same reason passing by value and calling setters
// will modify the original, e.g.:
//
//   func f1(val AttributeValue) { val.SetIntVal(234) }
//   func f2() {
//       v := NewAttributeValueString("a string")
//       f1(v)
//       _ := v.Type() // this will return AttributeValueTypeInt
//   }
//
// Important: zero-initialized instance is not valid for use. All AttributeValue functions below must
// be called only on instances that are created via NewAttributeValue+ functions.
type AttributeValue struct {
	pdata.AttributeValue
}

// NewAttributeValueEmpty creates a new AttributeValue with an empty value.
func NewAttributeValueEmpty() AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{})}
}

// NewAttributeValueString creates a new AttributeValue with the given string value.
func NewAttributeValueString(v string) AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_StringValue{StringValue: v}})}
}

// NewAttributeValueInt creates a new AttributeValue with the given int64 value.
func NewAttributeValueInt(v int64) AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_IntValue{IntValue: v}})}
}

// NewAttributeValueDouble creates a new AttributeValue with the given float64 value.
func NewAttributeValueDouble(v float64) AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_DoubleValue{DoubleValue: v}})}
}

// NewAttributeValueBool creates a new AttributeValue with the given bool value.
func NewAttributeValueBool(v bool) AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_BoolValue{BoolValue: v}})}
}

// NewAttributeValueMap creates a new AttributeValue of map type.
func NewAttributeValueMap() AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_KvlistValue{KvlistValue: &otlpcommon.KeyValueList{}}})}
}

// NewAttributeValueArray creates a new AttributeValue of array type.
func NewAttributeValueArray() AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_ArrayValue{ArrayValue: &otlpcommon.ArrayValue{}}})}
}

// NewAttributeValueBytes creates a new AttributeValue with the given []byte value.
// The caller must ensure the []byte passed in is not modified after the call is made, sharing the data
// across multiple attributes is forbidden.
func NewAttributeValueBytes(v []byte) AttributeValue {
	return AttributeValue{pdata.NewAttributeValue(&otlpcommon.AnyValue{Value: &otlpcommon.
		AnyValue_BytesValue{BytesValue: v}})}
}

// Type returns the type of the value for this AttributeValue.
// Calling this function on zero-initialized AttributeValue will cause a panic.
func (a AttributeValue) Type() AttributeValueType {
	return AttributeValueType(a.AttributeValue.Type())
}

// MapVal returns the map value associated with this AttributeValue.
// If the Type() is not AttributeValueTypeMap then returns an empty map. Note that modifying
// such empty map has no effect on this AttributeValue.
//
// Calling this function on zero-initialized AttributeValue will cause a panic.
func (a AttributeValue) MapVal() AttributeMap {
	return AttributeMap{a.AttributeValue.MapVal()}
}

// SliceVal returns the slice value associated with this AttributeValue.
// If the Type() is not AttributeValueTypeArray then returns an empty slice. Note that modifying
// such empty slice has no effect on this AttributeValue.
//
// Calling this function on zero-initialized AttributeValue will cause a panic.
func (a AttributeValue) SliceVal() AttributeValueSlice {
	return AttributeValueSlice{a.AttributeValue.SliceVal()}
}

// CopyTo copies the attribute to a destination.
func (a AttributeValue) CopyTo(dest AttributeValue) {
	a.AttributeValue.CopyTo(dest.AttributeValue)
}

// Equal checks for equality, it returns true if the objects are equal otherwise false.
func (a AttributeValue) Equal(av AttributeValue) bool {
	return a.AttributeValue.Equal(av.AttributeValue)
}

// AttributeMap stores a map of attribute keys to values.
type AttributeMap struct {
	pdata.AttributeMap
}

// NewAttributeMap creates a AttributeMap with 0 elements.
func NewAttributeMap() AttributeMap {
	orig := []otlpcommon.KeyValue(nil)
	return AttributeMap{pdata.NewAttributeMap(&orig)}
}

// NewAttributeMapFromMap creates a AttributeMap with values
// from the given map[string]AttributeValue.
func NewAttributeMapFromMap(attrMap map[string]AttributeValue) AttributeMap {
	unwrapped := make(map[string]pdata.AttributeValue, len(attrMap))
	for k, v := range attrMap {
		unwrapped[k] = v.AttributeValue
	}
	return AttributeMap{pdata.NewAttributeMapFromMap(unwrapped)}
}

// Get returns the AttributeValue associated with the key and true. Returned
// AttributeValue is not a copy, it is a reference to the value stored in this map.
// It is allowed to modify the returned value using AttributeValue.Set* functions.
// Such modification will be applied to the value stored in this map.
//
// If the key does not exist returns an invalid instance of the KeyValue and false.
// Calling any functions on the returned invalid instance will cause a panic.
func (am AttributeMap) Get(key string) (AttributeValue, bool) {
	v, ok := am.AttributeMap.Get(key)
	return AttributeValue{v}, ok
}

// Insert adds the AttributeValue to the map when the key does not exist.
// No action is applied to the map where the key already exists.
//
// Calling this function with a zero-initialized AttributeValue struct will cause a panic.
//
// Important: this function should not be used if the caller has access to
// the raw value to avoid an extra allocation.
func (am AttributeMap) Insert(k string, v AttributeValue) {
	am.AttributeMap.Insert(k, v.AttributeValue)
}

// Update updates an existing AttributeValue with a value.
// No action is applied to the map where the key does not exist.
//
// Calling this function with a zero-initialized AttributeValue struct will cause a panic.
//
// Important: this function should not be used if the caller has access to
// the raw value to avoid an extra allocation.
func (am AttributeMap) Update(k string, v AttributeValue) {
	am.AttributeMap.Update(k, v.AttributeValue)
}

// Upsert performs the Insert or Update action. The AttributeValue is
// inserted to the map that did not originally have the key. The key/value is
// updated to the map where the key already existed.
//
// Calling this function with a zero-initialized AttributeValue struct will cause a panic.
//
// Important: this function should not be used if the caller has access to
// the raw value to avoid an extra allocation.
func (am AttributeMap) Upsert(k string, v AttributeValue) {
	am.AttributeMap.Upsert(k, v.AttributeValue)
}

// Sort sorts the entries in the AttributeMap so two instances can be compared.
// Returns the same instance to allow nicer code like:
//   assert.EqualValues(t, expected.Sort(), actual.Sort())
func (am AttributeMap) Sort() AttributeMap {
	return AttributeMap{am.AttributeMap.Sort()}
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
//
// Example:
//
//   sm.Range(func(k string, v AttributeValue) bool {
//       ...
//   })
func (am AttributeMap) Range(f func(k string, v AttributeValue) bool) {
	am.AttributeMap.Range(func(k string, av pdata.AttributeValue) bool {
		return f(k, AttributeValue{av})
	})
}

// CopyTo copies all elements from the current map to the dest.
func (am AttributeMap) CopyTo(dest AttributeMap) {
	am.AttributeMap.CopyTo(dest.AttributeMap)
}
