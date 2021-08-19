// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_flag", Type: field.TypeBool, Default: false},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "food_id", Type: field.TypeUint64},
		{Name: "comment", Type: field.TypeString},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:        "comments",
		Columns:     CommentsColumns,
		PrimaryKey:  []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
	}
)

func init() {
}
