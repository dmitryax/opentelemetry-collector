// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pentity

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpentities "go.opentelemetry.io/collector/pdata/internal/data/protogen/entities/v1"
)

func TestScopeEntitiesSlice(t *testing.T) {
	es := NewScopeEntitiesSlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newScopeEntitiesSlice(&[]*otlpentities.ScopeEntities{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewScopeEntities()
	testVal := generateTestScopeEntities()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestScopeEntities(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestScopeEntitiesSliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newScopeEntitiesSlice(&[]*otlpentities.ScopeEntities{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewScopeEntitiesSlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestScopeEntitiesSlice_CopyTo(t *testing.T) {
	dest := NewScopeEntitiesSlice()
	// Test CopyTo to empty
	NewScopeEntitiesSlice().CopyTo(dest)
	assert.Equal(t, NewScopeEntitiesSlice(), dest)

	// Test CopyTo larger slice
	generateTestScopeEntitiesSlice().CopyTo(dest)
	assert.Equal(t, generateTestScopeEntitiesSlice(), dest)

	// Test CopyTo same size slice
	generateTestScopeEntitiesSlice().CopyTo(dest)
	assert.Equal(t, generateTestScopeEntitiesSlice(), dest)
}

func TestScopeEntitiesSlice_EnsureCapacity(t *testing.T) {
	es := generateTestScopeEntitiesSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestScopeEntitiesSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestScopeEntitiesSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestScopeEntitiesSlice(), es)
}

func TestScopeEntitiesSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestScopeEntitiesSlice()
	dest := NewScopeEntitiesSlice()
	src := generateTestScopeEntitiesSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestScopeEntitiesSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestScopeEntitiesSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestScopeEntitiesSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestScopeEntitiesSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewScopeEntitiesSlice()
	emptySlice.RemoveIf(func(el ScopeEntities) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestScopeEntitiesSlice()
	pos := 0
	filtered.RemoveIf(func(el ScopeEntities) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestScopeEntitiesSlice_Sort(t *testing.T) {
	es := generateTestScopeEntitiesSlice()
	es.Sort(func(a, b ScopeEntities) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Less(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b ScopeEntities) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Greater(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestScopeEntitiesSlice() ScopeEntitiesSlice {
	es := NewScopeEntitiesSlice()
	fillTestScopeEntitiesSlice(es)
	return es
}

func fillTestScopeEntitiesSlice(es ScopeEntitiesSlice) {
	*es.orig = make([]*otlpentities.ScopeEntities, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpentities.ScopeEntities{}
		fillTestScopeEntities(newScopeEntities((*es.orig)[i], es.state))
	}
}