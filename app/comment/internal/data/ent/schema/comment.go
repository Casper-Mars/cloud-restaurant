package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Casper-Mars/cloud-restaurant/pkg/myent"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return append(
		myent.BaseFields(),
		field.Uint64("user_id"),
		field.Uint64("food_id"),
		field.String("comment"),
	)
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
