package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/Casper-Mars/cloud-restaurant/pkg/myent"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return append(myent.BaseFields(),
		field.String("name"),
		field.String("phone").NotEmpty(),
		field.String("pwd").NotEmpty(),
	)
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
