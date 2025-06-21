package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SourceURL holds the schema definition for the SourceURL entity.
type SourceURL struct {
	ent.Schema
}

func (SourceURL) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (SourceURL) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

func (SourceURL) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sources", Source.Type),
	}
}
