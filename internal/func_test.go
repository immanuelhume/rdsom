package internal

import (
	"testing"
)

func TestChangeCase(t *testing.T) {
	fs := []struct {
		name string
		f    func(s string) string
	}{
		{
			name: "camelcase",
			f:    toCamelCase,
		},
		{
			name: "snakecase",
			f:    toSnakeCase,
		},
	}

	tests := []struct {
		ss       []string
		variants map[string]string
	}{
		{
			ss:       []string{"foo bar", "fooBar", "FooBar", "foo_bar"},
			variants: map[string]string{"camelcase": "fooBar", "snakecase": "foo_bar"},
		},
	}

	for _, f := range fs {
		for _, tt := range tests {
			for _, s := range tt.ss {
				got := f.f(s)
				want := tt.variants[f.name]
				if got != want {
					t.Errorf("got %q want %q when converting %q to %s", got, want, s, f.name)
				}
			}
		}
	}
}
