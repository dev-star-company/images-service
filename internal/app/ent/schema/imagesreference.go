package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ImagesReference holds the schema definition for the ImagesReference entity.
type ImagesReference struct {
	ent.Schema
}

func (ImagesReference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ImagesReference.
func (ImagesReference) Fields() []ent.Field {
	return []ent.Field{
		field.Int("image_id").Nillable(),
		field.Int("source_url_id").Optional().Nillable(),
		field.String("filename").NotEmpty(),
		field.String("mime_type").NotEmpty(),
		field.Bool("is_main").Default(false),
	}
}

// Edges of the ImagesReference.
func (ImagesReference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Images.Type),
		edge.To("source_url", SourceURL.Type),
	}
}
