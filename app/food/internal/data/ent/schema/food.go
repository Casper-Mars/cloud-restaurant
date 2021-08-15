package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Casper-Mars/cloud-restaurant/pkg/myent"
)

// Food holds the schema definition for the Food entity.
type Food struct {
	ent.Schema
}

// Fields of the Food.
func (Food) Fields() []ent.Field {
	return append(myent.BaseFields(),
		field.String("name"),
	)
}

// Edges of the Food.
func (Food) Edges() []ent.Edge {
	return nil
}
