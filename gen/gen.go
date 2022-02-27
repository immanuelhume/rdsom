package gen

import (
	"github.com/immanuelhume/rdsom"
)

func Gen(ss []rdsom.ISchema, metadata Metadata) error {
	return nil
}

type Metadata struct {
	Timestamp    uint64
	ProjName     string
	RdsomPkgPath string
}
