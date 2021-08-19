package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Food holds the schema definition for the Food entity.
type Food struct {
	ent.Schema
}

// Fields of the Food.
func (Food) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createTime").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updateTime").Default(func() time.Time {
			return time.Now()
		}).UpdateDefault(func() time.Time {
			return time.Now()
		}),
		field.String("name"),
	}
}

// Edges of the Food.
func (Food) Edges() []ent.Edge {
	return nil
}
