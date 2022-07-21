// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/revision"
	"ariga.io/atlas/cmd/atlas/internal/migrate/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	revisionFields := schema.Revision{}.Fields()
	_ = revisionFields
	// revisionDescTotal is the schema descriptor for total field.
	revisionDescTotal := revisionFields[3].Descriptor()
	// revision.TotalValidator is a validator for the "total" field. It is called by the builders before save.
	revision.TotalValidator = revisionDescTotal.Validators[0].(func(int) error)
}