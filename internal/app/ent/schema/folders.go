package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Folders holds the schema definition for the Folders entity.
type Folders struct {
	ent.Schema
}

func (Folders) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Folders.
func (Folders) Fields() []ent.Field {
	return []ent.Field{
		field.Int("folder_id").Optional().Nillable(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Folders.
func (Folders) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Images.Type),
		edge.To("children", Folders.Type).
			From("parent").
			Unique().
			Field("folder_id"),
	}
}
