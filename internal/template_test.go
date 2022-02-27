package internal_test

import (
	"bytes"
	"path"
	"testing"

	"github.com/immanuelhume/rdsom/internal"
	"github.com/immanuelhume/rdsom/internal/diff"
)

// TestAllTemplates is documentation for what data
// goes into each template. It allows for an simple,
// iterative approach to building out the template files.
func TestAllTemplates(t *testing.T) {
	tests := []struct {
		tmpl   string      // template name, e.g. "main.gotmpl"
		golden string      // name of the golden file
		data   interface{} // data to be passed into template
		noFmt  bool        // if the Go formatter should not be run
	}{

		{
			tmpl:   "main.gotmpl",
			golden: "_main.go",
			data: struct {
				SchemaPkgPath string
				Timestamp     uint64   // timestamp for this run
				SchemaNames   []string // schema names in alphabetical order
				ProjName      string
				RdsomPkgPath  string
			}{
				SchemaPkgPath: "github.com/foo/bar/schemas",
				Timestamp:     codeGenAt,
				SchemaNames:   []string{"Bar", "Foo"},
				ProjName:      "rdsomgolden",
				RdsomPkgPath:  "github.com/immanuelhume/rdsomgolden/rdsom",
			},
		},

		{
			tmpl:   "newschema.gotmpl",
			golden: "schema/_foo.go",
			data: struct {
				SchemaName string
			}{
				SchemaName: "foo",
			},
		},

		{
			tmpl:   "create.gotmpl",
			golden: "bar_create.go",
			data: struct {
				SchemaName   string
				Fields       []internal.Field
				ProjName     string
				RdsomPkgPath string
			}{
				SchemaName:   _barSchema.Name,
				Fields:       _barSchema.Fields,
				ProjName:     "rdsomgolden",
				RdsomPkgPath: "github.com/immanuelhume/rdsomgolden",
			},
		},

		{
			tmpl:   "schema.gotmpl",
			golden: "bar.go",
			data: struct {
				SchemaName string
				Fields     []internal.Field
				ProjName   string
				Timestamp  uint64
			}{
				SchemaName: _barSchema.Name,
				ProjName:   "rdsomgolden",
				Fields:     _barSchema.Fields,
				Timestamp:  codeGenAt,
			},
		},

		{
			tmpl:   "client.gotmpl",
			golden: "client.go",
			data: struct {
				Schemas []internal.Schema
			}{
				Schemas: []internal.Schema{_barSchema},
			},
		},

		{
			tmpl:   "predicate.gotmpl",
			golden: "predicate/predicate.go",
		},

		{
			tmpl:   "migrate.luatmpl",
			golden: "lua/migrate.lua",
			noFmt:  true,
			data: struct {
				ProjName  string
				Timestamp uint64
				Schemas   []internal.Schema
			}{
				ProjName:  "rdsomgolden",
				Timestamp: codeGenAt,
				Schemas:   []internal.Schema{_barSchema},
			},
		},
		{
			tmpl:   "where.gotmpl",
			golden: "bar/where.go",
			data: struct {
				Schema       internal.Schema
				RdsomPkgPath string
			}{
				Schema:       _barSchema,
				RdsomPkgPath: "github.com/immanuelhume/rdsomgolden",
			},
		},
	}

	for _, tt := range tests {
		var (
			golden    = path.Join("testdata", "golden", tt.golden)
			buf       = bytes.NewBuffer(nil)
			templater = internal.NewTemplater(buf, tt.tmpl)
		)
		if tt.noFmt {
			templater.SkipFmt()
		}
		err := templater.Do(tt.data)
		if err != nil {
			t.Fatal(err)
		}
		v, err := diff.NewDiff(buf, golden).Do()
		if err != nil {
			t.Fatal(err)
		}
		if v != nil {
			err := v.View()
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("template %q did not match golden %q", tt.tmpl, tt.golden)
		}
	}
}
