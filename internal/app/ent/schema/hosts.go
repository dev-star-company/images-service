package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Schema para representar os tipos de entidades que podem ter imagens vinculadas (ex: "product", "category")
type Hosts struct {
	ent.Schema
}

func (Hosts) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Hosts) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

func (Hosts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("host_urls", HostURLS.Type),
	}
}
