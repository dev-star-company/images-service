package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Images holds the schema definition for the Images entity.
type Images struct {
	ent.Schema
}

func (Images) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Images.
func (Images) Fields() []ent.Field {
	return []ent.Field{
		field.Int("media_type_id").Nillable().Optional(),
		field.Int("folder_id").Nillable().Optional(),
		field.String("name"),
		field.String("cloudflare_id").Optional(),
		field.String("url").Optional(),
		field.Int64("size").Optional(),
		field.String("content_type").Optional(),
	}
}

// Edges of the Images.
func (Images) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("folder", Folders.Type).
			Ref("images").
			Field("folder_id").
			Unique(),
		edge.From("media_type", MediaTypes.Type).
			Ref("images").
			Field("media_type_id").
			Unique(),
		edge.To("tags", Tags.Type),
	}
}
