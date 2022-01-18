package gen

import (
	"github.com/immanuelhume/rdsom"
	"github.com/immanuelhume/rdsom/internal"
)

func Gen(ss []rdsom.ISchema, metadata Metadata) error {
	return internal.StartCodeGen()
}

type Metadata struct {
	Timestamp    uint64
	ProjName     string
	RdsomPkgPath string
}
