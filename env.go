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

import "strings"

// tagOptions is the comma-separated options in a struct field's tag.
type tagOptions map[string]string

// Lookup retrieves the value of the option named by the key.
// If the option is present in the tag the value (which may be empty) is
// returned and the boolean is true. Otherwise, the returned value will be
// empty and the boolean will be false.
func (options tagOptions) Lookup(key string) (string, bool) {
	if options == nil {
		return "", false
	}

	value, exists := options[key]
	return value, exists
}

// cut slices s around the first unescaped instance of sep,
// returning the text before and after sep.
// The found result reports whether unescaped sep appears in s.
// If unescaped sep does not appear in s, cut returns s, "", false.
func cut(s, sep string) (before string, after string, found bool) {
	var escaping bool
	var afterBuilder strings.Builder
	var beforeBuilder strings.Builder

	for i := 0; i < len(s); i++ {
		if !escaping && s[i] == '\\' {
			escaping = true
			continue
		}

		if !escaping && i <= len(s)-len(sep) && strings.HasPrefix(s[i:], sep) {
			for _, c := range s[i+len(sep):] {
				if !escaping && c == '\\' {
					escaping = true
					continue
				}

				afterBuilder.WriteRune(c)
				escaping = false
			}

			return beforeBuilder.String(), afterBuilder.String(), true
		}

		beforeBuilder.WriteByte(s[i])
		escaping = false
	}

	return beforeBuilder.String(), "", false
}
