package env

import (
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		sep        string
		wantBefore string
		wantAfter  string
		wantFound  bool
	}{
		{"empty string", "", ",", "", "", false},
		{"no separator", "abc", ",", "abc", "", false},
		{"empty separator", "abc", "", "", "abc", true},
		{"basic separator", "a,b,c", ",", "a", "b,c", true},
		{"unicode separator", "a≠b≠c", "≠", "a", "b≠c", true},
		{"multi-byte separator", "a::b::c", "::", "a", "b::c", true},

		{"escaped separator", `a\,b,c`, ",", "a,b", "c", true},
		{"escaped unicode separator", `a\≠b≠c`, "≠", `a≠b`, "c", true},
		{"escaped multi-byte separator", `a\::b::c`, "::", "a::b", "c", true},
		{"all separators escaped", `a\,b\,c`, ",", "a,b,c", "", false},
		{"backslash before separator", `a\\,b`, ",", `a\`, "b", true},
		{"multiple backslashes", `a\\\,b`, ",", `a\,b`, "", false},

		{"separator at start", ",a,b", ",", "", "a,b", true},
		{"separator at end", "a,b,", ",", "a", "b,", true},
		{"only separator", ",", ",", "", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before, after, found := cut(test.s, test.sep)

			if before != test.wantBefore {
				t.Errorf("before: got %q, want %q", before, test.wantBefore)
			}

			if after != test.wantAfter {
				t.Errorf("after: got %q, want %q", after, test.wantAfter)
			}

			if found != test.wantFound {
				t.Errorf("found: got %v, want %v", found, test.wantFound)
			}
		})
	}
}
