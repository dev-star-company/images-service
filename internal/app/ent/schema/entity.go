package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Schema para representar os tipos de entidades que podem ter imagens vinculadas (ex: "product", "category")
type Entity struct {
	ent.Schema
}

func (Entity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

func (Entity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Images.Type),
	}
}
