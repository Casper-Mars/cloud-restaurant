package myent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/Casper-Mars/cloud-restaurant/pkg/enum"
	"time"
)

func BaseFields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Annotations(entsql.Annotation{
			Incremental: &enum.Incremental,
		}),
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
