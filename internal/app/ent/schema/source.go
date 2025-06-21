package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Source struct {
	ent.Schema
}

func (Source) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}

}

func (Source) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source_url", SourceURL.Type).
			Ref("sources").
			Unique().
			Required(),
	}
}
