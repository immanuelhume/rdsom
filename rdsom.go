package rdsom

import "github.com/immanuelhume/rdsom/internal"

// Schema should be embedded inside a user-defined schema like this.
//
//	type Foo struct {
//		rdsom.Schema
//	}
//
// This allows the compiler accepts Foo as an ISchema.
type Schema struct {
	ISchema
}

func (s *Schema) Fields() []Field {
	return nil
}

func (s *Schema) Edges() []Edge {
	return nil
}

// ISchema is an abstraction of a user-defined schema for use in the rdsom program.
type ISchema interface {
	Fields() []Field
	Edges() []Edge
}

type Field interface {
	Meta() *internal.Field
}

type Edge interface {
	Meta() *internal.Edge
}
