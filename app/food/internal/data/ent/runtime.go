// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/data/ent/food"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	foodFields := schema.Food{}.Fields()
	_ = foodFields
	// foodDescCreateTime is the schema descriptor for createTime field.
	foodDescCreateTime := foodFields[1].Descriptor()
	// food.DefaultCreateTime holds the default value on creation for the createTime field.
	food.DefaultCreateTime = foodDescCreateTime.Default.(func() time.Time)
	// foodDescUpdateTime is the schema descriptor for updateTime field.
	foodDescUpdateTime := foodFields[2].Descriptor()
	// food.DefaultUpdateTime holds the default value on creation for the updateTime field.
	food.DefaultUpdateTime = foodDescUpdateTime.Default.(func() time.Time)
	// food.UpdateDefaultUpdateTime holds the default value on update for the updateTime field.
	food.UpdateDefaultUpdateTime = foodDescUpdateTime.UpdateDefault.(func() time.Time)
	// foodDescDeleteFlag is the schema descriptor for delete_flag field.
	foodDescDeleteFlag := foodFields[3].Descriptor()
	// food.DefaultDeleteFlag holds the default value on creation for the delete_flag field.
	food.DefaultDeleteFlag = foodDescDeleteFlag.Default.(bool)
}
