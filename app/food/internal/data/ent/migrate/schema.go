// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FoodsColumns holds the columns for the "foods" table.
	FoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_flag", Type: field.TypeBool, Default: false},
		{Name: "name", Type: field.TypeString},
	}
	// FoodsTable holds the schema information for the "foods" table.
	FoodsTable = &schema.Table{
		Name:        "foods",
		Columns:     FoodsColumns,
		PrimaryKey:  []*schema.Column{FoodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FoodsTable,
	}
)

func init() {
}
