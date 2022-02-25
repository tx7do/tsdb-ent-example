// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"tsdb-ent-example/ent/attributekv"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AttributeKV is the model entity for the AttributeKV schema.
type AttributeKV struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EntityType holds the value of the "entity_type" field.
	// 实体类型
	EntityType string `json:"entity_type,omitempty"`
	// EntityID holds the value of the "entity_id" field.
	// 实体ID
	EntityID uuid.UUID `json:"entity_id,omitempty"`
	// AttributeType holds the value of the "attribute_type" field.
	// 属性类型
	AttributeType string `json:"attribute_type,omitempty"`
	// AttributeKey holds the value of the "attribute_key" field.
	// 属性键名
	AttributeKey string `json:"attribute_key,omitempty"`
	// BoolV holds the value of the "bool_v" field.
	// BOOL值
	BoolV *bool `json:"bool_v,omitempty"`
	// StrV holds the value of the "str_v" field.
	// STRING值
	StrV *string `json:"str_v,omitempty"`
	// LongV holds the value of the "long_v" field.
	// LONG值
	LongV *int64 `json:"long_v,omitempty"`
	// DblV holds the value of the "dbl_v" field.
	// DOUBLE值
	DblV *float64 `json:"dbl_v,omitempty"`
	// JSONV holds the value of the "json_v" field.
	// JSON值
	JSONV *string `json:"json_v,omitempty"`
	// LastUpdateTs holds the value of the "last_update_ts" field.
	// 最后更新的时间戳
	LastUpdateTs int64 `json:"last_update_ts,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeKV) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributekv.FieldBoolV:
			values[i] = new(sql.NullBool)
		case attributekv.FieldDblV:
			values[i] = new(sql.NullFloat64)
		case attributekv.FieldID, attributekv.FieldLongV, attributekv.FieldLastUpdateTs:
			values[i] = new(sql.NullInt64)
		case attributekv.FieldEntityType, attributekv.FieldAttributeType, attributekv.FieldAttributeKey, attributekv.FieldStrV, attributekv.FieldJSONV:
			values[i] = new(sql.NullString)
		case attributekv.FieldEntityID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AttributeKV", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeKV fields.
func (ak *AttributeKV) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributekv.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ak.ID = int(value.Int64)
		case attributekv.FieldEntityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_type", values[i])
			} else if value.Valid {
				ak.EntityType = value.String
			}
		case attributekv.FieldEntityID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field entity_id", values[i])
			} else if value != nil {
				ak.EntityID = *value
			}
		case attributekv.FieldAttributeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_type", values[i])
			} else if value.Valid {
				ak.AttributeType = value.String
			}
		case attributekv.FieldAttributeKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attribute_key", values[i])
			} else if value.Valid {
				ak.AttributeKey = value.String
			}
		case attributekv.FieldBoolV:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field bool_v", values[i])
			} else if value.Valid {
				ak.BoolV = new(bool)
				*ak.BoolV = value.Bool
			}
		case attributekv.FieldStrV:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field str_v", values[i])
			} else if value.Valid {
				ak.StrV = new(string)
				*ak.StrV = value.String
			}
		case attributekv.FieldLongV:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field long_v", values[i])
			} else if value.Valid {
				ak.LongV = new(int64)
				*ak.LongV = value.Int64
			}
		case attributekv.FieldDblV:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field dbl_v", values[i])
			} else if value.Valid {
				ak.DblV = new(float64)
				*ak.DblV = value.Float64
			}
		case attributekv.FieldJSONV:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field json_v", values[i])
			} else if value.Valid {
				ak.JSONV = new(string)
				*ak.JSONV = value.String
			}
		case attributekv.FieldLastUpdateTs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_update_ts", values[i])
			} else if value.Valid {
				ak.LastUpdateTs = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AttributeKV.
// Note that you need to call AttributeKV.Unwrap() before calling this method if this AttributeKV
// was returned from a transaction, and the transaction was committed or rolled back.
func (ak *AttributeKV) Update() *AttributeKVUpdateOne {
	return (&AttributeKVClient{config: ak.config}).UpdateOne(ak)
}

// Unwrap unwraps the AttributeKV entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ak *AttributeKV) Unwrap() *AttributeKV {
	tx, ok := ak.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeKV is not a transactional entity")
	}
	ak.config.driver = tx.drv
	return ak
}

// String implements the fmt.Stringer.
func (ak *AttributeKV) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeKV(")
	builder.WriteString(fmt.Sprintf("id=%v", ak.ID))
	builder.WriteString(", entity_type=")
	builder.WriteString(ak.EntityType)
	builder.WriteString(", entity_id=")
	builder.WriteString(fmt.Sprintf("%v", ak.EntityID))
	builder.WriteString(", attribute_type=")
	builder.WriteString(ak.AttributeType)
	builder.WriteString(", attribute_key=")
	builder.WriteString(ak.AttributeKey)
	if v := ak.BoolV; v != nil {
		builder.WriteString(", bool_v=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ak.StrV; v != nil {
		builder.WriteString(", str_v=")
		builder.WriteString(*v)
	}
	if v := ak.LongV; v != nil {
		builder.WriteString(", long_v=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ak.DblV; v != nil {
		builder.WriteString(", dbl_v=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ak.JSONV; v != nil {
		builder.WriteString(", json_v=")
		builder.WriteString(*v)
	}
	builder.WriteString(", last_update_ts=")
	builder.WriteString(fmt.Sprintf("%v", ak.LastUpdateTs))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeKVs is a parsable slice of AttributeKV.
type AttributeKVs []*AttributeKV

func (ak AttributeKVs) config(cfg config) {
	for _i := range ak {
		ak[_i].config = cfg
	}
}