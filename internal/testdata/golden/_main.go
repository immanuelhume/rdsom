package main

import (
	"log"

	"github.com/foo/bar/schemas"
	"github.com/immanuelhume/rdsom"
	"github.com/immanuelhume/rdsom/gen"
)

const _timestamp uint64 = 1642524233

var _bar = schemas.Bar{}
var _foo = schemas.Foo{}
var _schemas = []rdsom.ISchema{_foo, _bar}

func main() {
	metadata := gen.Metadata{
		Timestamp:    1642524233,
		ProjName:     "rdsomgolden",
		RdsomPkgPath: "github.com/immanuelhume/rdsomgolden/rdsom",
	}
	if err := gen.Gen(_schemas, metadata); err != nil {
		log.Fatal("running codegen: %s", err)
	}
}
