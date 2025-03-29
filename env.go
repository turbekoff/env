// Copyright 2025 Dias Turdybek. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package env implements encoding and decoding of environment variables.
// The mapping between environment variables and Golang values is
// described in the documentation for the Encode and Decode functions.
//
// See GitHub repository for an introduction to this package:
// https://github.com/turbekoff/env
package env

// tagOptions is the comma-separated options in a struct field's tag.
type tagOptions map[string]string
