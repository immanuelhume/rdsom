package internal

import (
	"fmt"
	"strings"
)

type Field struct {
	Name           string   // name of the field as defined on the struct type
	JsonName       string   // name of the field as defined by the user
	GoType         string   // string representation of the field's type in Go
	RediSearchOpts []string // options to pass FT.CREATE
}

// LuaParams is used for templating. It returns a string
// of options for use in the Lua migration script.
// e.g. `"FooField", "TEXT"`
func (f *Field) LuaParams() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("%q", f.JsonName))
	for _, lf := range f.RediSearchOpts {
		s.WriteString(fmt.Sprintf(", %q", lf))
	}
	return s.String()
}
