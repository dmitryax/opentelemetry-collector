// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pentity

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpentities "go.opentelemetry.io/collector/pdata/internal/data/protogen/entities/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceEntities is a collection of entities from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceEntities function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceEntities struct {
	orig  *otlpentities.ResourceEntities
	state *internal.State
}

func newResourceEntities(orig *otlpentities.ResourceEntities, state *internal.State) ResourceEntities {
	return ResourceEntities{orig: orig, state: state}
}

// NewResourceEntities creates a new empty ResourceEntities.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewResourceEntities() ResourceEntities {
	state := internal.StateMutable
	return newResourceEntities(&otlpentities.ResourceEntities{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ResourceEntities) MoveTo(dest ResourceEntities) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpentities.ResourceEntities{}
}

// Resource returns the resource associated with this ResourceEntities.
func (ms ResourceEntities) Resource() pcommon.Resource {
	return pcommon.Resource(internal.NewResource(&ms.orig.Resource, ms.state))
}

// SchemaUrl returns the schemaurl associated with this ResourceEntities.
func (ms ResourceEntities) SchemaUrl() string {
	return ms.orig.SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ResourceEntities.
func (ms ResourceEntities) SetSchemaUrl(v string) {
	ms.state.AssertMutable()
	ms.orig.SchemaUrl = v
}

// ScopeEntities returns the ScopeEntities associated with this ResourceEntities.
func (ms ResourceEntities) ScopeEntities() ScopeEntitiesSlice {
	return newScopeEntitiesSlice(&ms.orig.ScopeEntities, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ResourceEntities) CopyTo(dest ResourceEntities) {
	dest.state.AssertMutable()
	ms.Resource().CopyTo(dest.Resource())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.ScopeEntities().CopyTo(dest.ScopeEntities())
}