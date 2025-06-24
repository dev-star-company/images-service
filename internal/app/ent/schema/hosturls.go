package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HostURLS struct {
	ent.Schema
}

func (HostURLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (HostURLS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("hosts_id").Nillable(),
		field.String("name").NotEmpty(),
		field.String("url").NotEmpty(),
		field.Bool("default"),
	}

}

func (HostURLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host", Hosts.Type).
			Ref("host_urls").
			Unique(),
	}
}
