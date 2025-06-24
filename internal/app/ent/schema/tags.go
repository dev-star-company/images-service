package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Schema para representar os tipos de entidades que podem ter imagens vinculadas (ex: "product", "category")
type Tags struct {
	ent.Schema
}

func (Tags) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Tags) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

func (Tags) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("images", Images.Type).
			Ref("tags"),
	}
}
