package internal

import (
	"reflect"
	"testing"
)

func TestSchemaNames(t *testing.T) {
	cases := []struct {
		desc    string
		pkgPath string
		want    []string
	}{
		{
			desc:    "normal schema declarations",
			pkgPath: "github.com/immanuelhume/rdsom/internal/testdata/pkga/rdsom/schemas",
			want:    []string{"Bar", "Foo"},
		},
		{
			desc:    "no declarations",
			pkgPath: "github.com/immanuelhume/rdsom/internal/testdata/pkgb/rdsom/schemas",
			want:    nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := schemas(tt.pkgPath)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}
