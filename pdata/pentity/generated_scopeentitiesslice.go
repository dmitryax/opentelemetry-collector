// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pentity

import (
	"sort"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpentities "go.opentelemetry.io/collector/pdata/internal/data/protogen/entities/v1"
)

// ScopeEntitiesSlice logically represents a slice of ScopeEntities.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeEntitiesSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeEntitiesSlice struct {
	orig  *[]*otlpentities.ScopeEntities
	state *internal.State
}

func newScopeEntitiesSlice(orig *[]*otlpentities.ScopeEntities, state *internal.State) ScopeEntitiesSlice {
	return ScopeEntitiesSlice{orig: orig, state: state}
}

// NewScopeEntitiesSlice creates a ScopeEntitiesSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewScopeEntitiesSlice() ScopeEntitiesSlice {
	orig := []*otlpentities.ScopeEntities(nil)
	state := internal.StateMutable
	return newScopeEntitiesSlice(&orig, &state)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewScopeEntitiesSlice()".
func (es ScopeEntitiesSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ScopeEntitiesSlice) At(i int) ScopeEntities {
	return newScopeEntities((*es.orig)[i], es.state)
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ScopeEntitiesSlice can be initialized:
//
//	es := NewScopeEntitiesSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ScopeEntitiesSlice) EnsureCapacity(newCap int) {
	es.state.AssertMutable()
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlpentities.ScopeEntities, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty ScopeEntities.
// It returns the newly added ScopeEntities.
func (es ScopeEntitiesSlice) AppendEmpty() ScopeEntities {
	es.state.AssertMutable()
	*es.orig = append(*es.orig, &otlpentities.ScopeEntities{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ScopeEntitiesSlice) MoveAndAppendTo(dest ScopeEntitiesSlice) {
	es.state.AssertMutable()
	dest.state.AssertMutable()
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ScopeEntitiesSlice) RemoveIf(f func(ScopeEntities) bool) {
	es.state.AssertMutable()
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		newLen++
	}
	*es.orig = (*es.orig)[:newLen]
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ScopeEntitiesSlice) CopyTo(dest ScopeEntitiesSlice) {
	dest.state.AssertMutable()
	srcLen := es.Len()
	destCap := cap(*dest.orig)
	if srcLen <= destCap {
		(*dest.orig) = (*dest.orig)[:srcLen:destCap]
		for i := range *es.orig {
			newScopeEntities((*es.orig)[i], es.state).CopyTo(newScopeEntities((*dest.orig)[i], dest.state))
		}
		return
	}
	origs := make([]otlpentities.ScopeEntities, srcLen)
	wrappers := make([]*otlpentities.ScopeEntities, srcLen)
	for i := range *es.orig {
		wrappers[i] = &origs[i]
		newScopeEntities((*es.orig)[i], es.state).CopyTo(newScopeEntities(wrappers[i], dest.state))
	}
	*dest.orig = wrappers
}

// Sort sorts the ScopeEntities elements within ScopeEntitiesSlice given the
// provided less function so that two instances of ScopeEntitiesSlice
// can be compared.
func (es ScopeEntitiesSlice) Sort(less func(a, b ScopeEntities) bool) {
	es.state.AssertMutable()
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}