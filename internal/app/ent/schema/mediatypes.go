package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Schema para representar os tipos de entidades que podem ter imagens vinculadas (ex: "product", "category")
type MediaTypes struct {
	ent.Schema
}

func (MediaTypes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MediaTypes) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

func (MediaTypes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Images.Type),
	}
}
