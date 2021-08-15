// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Casper-Mars/cloud-restaurant/app/food/internal/data/ent/food"
)

// Food is the model entity for the Food schema.
type Food struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "createTime" field.
	CreateTime time.Time `json:"createTime,omitempty"`
	// UpdateTime holds the value of the "updateTime" field.
	UpdateTime time.Time `json:"updateTime,omitempty"`
	// DeleteFlag holds the value of the "delete_flag" field.
	DeleteFlag bool `json:"delete_flag,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Food) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case food.FieldDeleteFlag:
			values[i] = new(sql.NullBool)
		case food.FieldID:
			values[i] = new(sql.NullInt64)
		case food.FieldName:
			values[i] = new(sql.NullString)
		case food.FieldCreateTime, food.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Food", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Food fields.
func (f *Food) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case food.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = uint64(value.Int64)
		case food.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createTime", values[i])
			} else if value.Valid {
				f.CreateTime = value.Time
			}
		case food.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updateTime", values[i])
			} else if value.Valid {
				f.UpdateTime = value.Time
			}
		case food.FieldDeleteFlag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field delete_flag", values[i])
			} else if value.Valid {
				f.DeleteFlag = value.Bool
			}
		case food.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Food.
// Note that you need to call Food.Unwrap() before calling this method if this Food
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Food) Update() *FoodUpdateOne {
	return (&FoodClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Food entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Food) Unwrap() *Food {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Food is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Food) String() string {
	var builder strings.Builder
	builder.WriteString("Food(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", createTime=")
	builder.WriteString(f.CreateTime.Format(time.ANSIC))
	builder.WriteString(", updateTime=")
	builder.WriteString(f.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", delete_flag=")
	builder.WriteString(fmt.Sprintf("%v", f.DeleteFlag))
	builder.WriteString(", name=")
	builder.WriteString(f.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Foods is a parsable slice of Food.
type Foods []*Food

func (f Foods) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
