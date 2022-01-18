package predicate_test

import (
	"testing"

	"github.com/immanuelhume/rdsomgolden/predicate"
)

func TestEscapePredicate(t *testing.T) {
	tests := []struct {
		raw     string
		escaped string
	}{
		{raw: "foobar", escaped: "foobar"},
		{raw: "foo.bar", escaped: "foo\\.bar"},
		{raw: "[1,2.1,3.2]", escaped: "\\[1\\,2\\.1\\,3\\.2\\]"},
	}
	for _, tt := range tests {
		t.Run(tt.raw, func(t *testing.T) {
			// Tests escaping first.
			got := predicate.Escape(tt.raw)
			if got != tt.escaped {
				t.Errorf("got %q want %q", got, tt.escaped)
			}

			// Test unescaping.
			got = predicate.Unescape(tt.escaped)
			if got != tt.raw {
				t.Errorf("got %q want %q", got, tt.raw)
			}
		})
	}
}
