package schema

import (
	"entgo.io/ent"
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
		field.Int("entity_id"),
		field.Bool("is_main").Default(false),
	}
}

// Edges of the Images.
func (Images) Edges() []ent.Edge {
	return nil

}
