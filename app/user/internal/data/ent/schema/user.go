package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createTime").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updateTime").Default(func() time.Time {
			return time.Now()
		}).UpdateDefault(func() time.Time {
			return time.Now()
		}),
		field.Bool("delete_flag").Default(false),
		field.String("name"),
		field.String("phone").NotEmpty(),
		field.String("pwd").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
