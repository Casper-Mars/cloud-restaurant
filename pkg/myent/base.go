package myent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

func BaseFields() []ent.Field {
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
	}
}
