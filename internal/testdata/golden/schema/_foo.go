package schema

import "github.com/immanuelhume/rdsom"

type Foo struct {
	rdsom.Schema
}

func (f *Foo) Fields() []rdsom.Field {
	return []rdsom.Field{}
}
