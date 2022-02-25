// Code generated by entc, DO NOT EDIT.

package attributekv

import (
	"tsdb-ent-example/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EntityType applies equality check predicate on the "entity_type" field. It's identical to EntityTypeEQ.
func EntityType(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntityType), v))
	})
}

// EntityID applies equality check predicate on the "entity_id" field. It's identical to EntityIDEQ.
func EntityID(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntityID), v))
	})
}

// AttributeType applies equality check predicate on the "attribute_type" field. It's identical to AttributeTypeEQ.
func AttributeType(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttributeType), v))
	})
}

// AttributeKey applies equality check predicate on the "attribute_key" field. It's identical to AttributeKeyEQ.
func AttributeKey(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttributeKey), v))
	})
}

// BoolV applies equality check predicate on the "bool_v" field. It's identical to BoolVEQ.
func BoolV(v bool) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBoolV), v))
	})
}

// StrV applies equality check predicate on the "str_v" field. It's identical to StrVEQ.
func StrV(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrV), v))
	})
}

// LongV applies equality check predicate on the "long_v" field. It's identical to LongVEQ.
func LongV(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongV), v))
	})
}

// DblV applies equality check predicate on the "dbl_v" field. It's identical to DblVEQ.
func DblV(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDblV), v))
	})
}

// JSONV applies equality check predicate on the "json_v" field. It's identical to JSONVEQ.
func JSONV(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJSONV), v))
	})
}

// LastUpdateTs applies equality check predicate on the "last_update_ts" field. It's identical to LastUpdateTsEQ.
func LastUpdateTs(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdateTs), v))
	})
}

// EntityTypeEQ applies the EQ predicate on the "entity_type" field.
func EntityTypeEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntityType), v))
	})
}

// EntityTypeNEQ applies the NEQ predicate on the "entity_type" field.
func EntityTypeNEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntityType), v))
	})
}

// EntityTypeIn applies the In predicate on the "entity_type" field.
func EntityTypeIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntityType), v...))
	})
}

// EntityTypeNotIn applies the NotIn predicate on the "entity_type" field.
func EntityTypeNotIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntityType), v...))
	})
}

// EntityTypeGT applies the GT predicate on the "entity_type" field.
func EntityTypeGT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntityType), v))
	})
}

// EntityTypeGTE applies the GTE predicate on the "entity_type" field.
func EntityTypeGTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntityType), v))
	})
}

// EntityTypeLT applies the LT predicate on the "entity_type" field.
func EntityTypeLT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntityType), v))
	})
}

// EntityTypeLTE applies the LTE predicate on the "entity_type" field.
func EntityTypeLTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntityType), v))
	})
}

// EntityTypeContains applies the Contains predicate on the "entity_type" field.
func EntityTypeContains(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEntityType), v))
	})
}

// EntityTypeHasPrefix applies the HasPrefix predicate on the "entity_type" field.
func EntityTypeHasPrefix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEntityType), v))
	})
}

// EntityTypeHasSuffix applies the HasSuffix predicate on the "entity_type" field.
func EntityTypeHasSuffix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEntityType), v))
	})
}

// EntityTypeEqualFold applies the EqualFold predicate on the "entity_type" field.
func EntityTypeEqualFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEntityType), v))
	})
}

// EntityTypeContainsFold applies the ContainsFold predicate on the "entity_type" field.
func EntityTypeContainsFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEntityType), v))
	})
}

// EntityIDEQ applies the EQ predicate on the "entity_id" field.
func EntityIDEQ(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntityID), v))
	})
}

// EntityIDNEQ applies the NEQ predicate on the "entity_id" field.
func EntityIDNEQ(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntityID), v))
	})
}

// EntityIDIn applies the In predicate on the "entity_id" field.
func EntityIDIn(vs ...uuid.UUID) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntityID), v...))
	})
}

// EntityIDNotIn applies the NotIn predicate on the "entity_id" field.
func EntityIDNotIn(vs ...uuid.UUID) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntityID), v...))
	})
}

// EntityIDGT applies the GT predicate on the "entity_id" field.
func EntityIDGT(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntityID), v))
	})
}

// EntityIDGTE applies the GTE predicate on the "entity_id" field.
func EntityIDGTE(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntityID), v))
	})
}

// EntityIDLT applies the LT predicate on the "entity_id" field.
func EntityIDLT(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntityID), v))
	})
}

// EntityIDLTE applies the LTE predicate on the "entity_id" field.
func EntityIDLTE(v uuid.UUID) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntityID), v))
	})
}

// AttributeTypeEQ applies the EQ predicate on the "attribute_type" field.
func AttributeTypeEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeNEQ applies the NEQ predicate on the "attribute_type" field.
func AttributeTypeNEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeIn applies the In predicate on the "attribute_type" field.
func AttributeTypeIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttributeType), v...))
	})
}

// AttributeTypeNotIn applies the NotIn predicate on the "attribute_type" field.
func AttributeTypeNotIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttributeType), v...))
	})
}

// AttributeTypeGT applies the GT predicate on the "attribute_type" field.
func AttributeTypeGT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeGTE applies the GTE predicate on the "attribute_type" field.
func AttributeTypeGTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeLT applies the LT predicate on the "attribute_type" field.
func AttributeTypeLT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeLTE applies the LTE predicate on the "attribute_type" field.
func AttributeTypeLTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeContains applies the Contains predicate on the "attribute_type" field.
func AttributeTypeContains(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeHasPrefix applies the HasPrefix predicate on the "attribute_type" field.
func AttributeTypeHasPrefix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeHasSuffix applies the HasSuffix predicate on the "attribute_type" field.
func AttributeTypeHasSuffix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeEqualFold applies the EqualFold predicate on the "attribute_type" field.
func AttributeTypeEqualFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAttributeType), v))
	})
}

// AttributeTypeContainsFold applies the ContainsFold predicate on the "attribute_type" field.
func AttributeTypeContainsFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAttributeType), v))
	})
}

// AttributeKeyEQ applies the EQ predicate on the "attribute_key" field.
func AttributeKeyEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyNEQ applies the NEQ predicate on the "attribute_key" field.
func AttributeKeyNEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyIn applies the In predicate on the "attribute_key" field.
func AttributeKeyIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttributeKey), v...))
	})
}

// AttributeKeyNotIn applies the NotIn predicate on the "attribute_key" field.
func AttributeKeyNotIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttributeKey), v...))
	})
}

// AttributeKeyGT applies the GT predicate on the "attribute_key" field.
func AttributeKeyGT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyGTE applies the GTE predicate on the "attribute_key" field.
func AttributeKeyGTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyLT applies the LT predicate on the "attribute_key" field.
func AttributeKeyLT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyLTE applies the LTE predicate on the "attribute_key" field.
func AttributeKeyLTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyContains applies the Contains predicate on the "attribute_key" field.
func AttributeKeyContains(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyHasPrefix applies the HasPrefix predicate on the "attribute_key" field.
func AttributeKeyHasPrefix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyHasSuffix applies the HasSuffix predicate on the "attribute_key" field.
func AttributeKeyHasSuffix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyEqualFold applies the EqualFold predicate on the "attribute_key" field.
func AttributeKeyEqualFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAttributeKey), v))
	})
}

// AttributeKeyContainsFold applies the ContainsFold predicate on the "attribute_key" field.
func AttributeKeyContainsFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAttributeKey), v))
	})
}

// BoolVEQ applies the EQ predicate on the "bool_v" field.
func BoolVEQ(v bool) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBoolV), v))
	})
}

// BoolVNEQ applies the NEQ predicate on the "bool_v" field.
func BoolVNEQ(v bool) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBoolV), v))
	})
}

// BoolVIsNil applies the IsNil predicate on the "bool_v" field.
func BoolVIsNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBoolV)))
	})
}

// BoolVNotNil applies the NotNil predicate on the "bool_v" field.
func BoolVNotNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBoolV)))
	})
}

// StrVEQ applies the EQ predicate on the "str_v" field.
func StrVEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrV), v))
	})
}

// StrVNEQ applies the NEQ predicate on the "str_v" field.
func StrVNEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStrV), v))
	})
}

// StrVIn applies the In predicate on the "str_v" field.
func StrVIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStrV), v...))
	})
}

// StrVNotIn applies the NotIn predicate on the "str_v" field.
func StrVNotIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStrV), v...))
	})
}

// StrVGT applies the GT predicate on the "str_v" field.
func StrVGT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStrV), v))
	})
}

// StrVGTE applies the GTE predicate on the "str_v" field.
func StrVGTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStrV), v))
	})
}

// StrVLT applies the LT predicate on the "str_v" field.
func StrVLT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStrV), v))
	})
}

// StrVLTE applies the LTE predicate on the "str_v" field.
func StrVLTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStrV), v))
	})
}

// StrVContains applies the Contains predicate on the "str_v" field.
func StrVContains(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStrV), v))
	})
}

// StrVHasPrefix applies the HasPrefix predicate on the "str_v" field.
func StrVHasPrefix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStrV), v))
	})
}

// StrVHasSuffix applies the HasSuffix predicate on the "str_v" field.
func StrVHasSuffix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStrV), v))
	})
}

// StrVIsNil applies the IsNil predicate on the "str_v" field.
func StrVIsNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStrV)))
	})
}

// StrVNotNil applies the NotNil predicate on the "str_v" field.
func StrVNotNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStrV)))
	})
}

// StrVEqualFold applies the EqualFold predicate on the "str_v" field.
func StrVEqualFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStrV), v))
	})
}

// StrVContainsFold applies the ContainsFold predicate on the "str_v" field.
func StrVContainsFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStrV), v))
	})
}

// LongVEQ applies the EQ predicate on the "long_v" field.
func LongVEQ(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongV), v))
	})
}

// LongVNEQ applies the NEQ predicate on the "long_v" field.
func LongVNEQ(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongV), v))
	})
}

// LongVIn applies the In predicate on the "long_v" field.
func LongVIn(vs ...int64) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLongV), v...))
	})
}

// LongVNotIn applies the NotIn predicate on the "long_v" field.
func LongVNotIn(vs ...int64) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLongV), v...))
	})
}

// LongVGT applies the GT predicate on the "long_v" field.
func LongVGT(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongV), v))
	})
}

// LongVGTE applies the GTE predicate on the "long_v" field.
func LongVGTE(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongV), v))
	})
}

// LongVLT applies the LT predicate on the "long_v" field.
func LongVLT(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongV), v))
	})
}

// LongVLTE applies the LTE predicate on the "long_v" field.
func LongVLTE(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongV), v))
	})
}

// LongVIsNil applies the IsNil predicate on the "long_v" field.
func LongVIsNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLongV)))
	})
}

// LongVNotNil applies the NotNil predicate on the "long_v" field.
func LongVNotNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLongV)))
	})
}

// DblVEQ applies the EQ predicate on the "dbl_v" field.
func DblVEQ(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDblV), v))
	})
}

// DblVNEQ applies the NEQ predicate on the "dbl_v" field.
func DblVNEQ(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDblV), v))
	})
}

// DblVIn applies the In predicate on the "dbl_v" field.
func DblVIn(vs ...float64) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDblV), v...))
	})
}

// DblVNotIn applies the NotIn predicate on the "dbl_v" field.
func DblVNotIn(vs ...float64) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDblV), v...))
	})
}

// DblVGT applies the GT predicate on the "dbl_v" field.
func DblVGT(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDblV), v))
	})
}

// DblVGTE applies the GTE predicate on the "dbl_v" field.
func DblVGTE(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDblV), v))
	})
}

// DblVLT applies the LT predicate on the "dbl_v" field.
func DblVLT(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDblV), v))
	})
}

// DblVLTE applies the LTE predicate on the "dbl_v" field.
func DblVLTE(v float64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDblV), v))
	})
}

// DblVIsNil applies the IsNil predicate on the "dbl_v" field.
func DblVIsNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDblV)))
	})
}

// DblVNotNil applies the NotNil predicate on the "dbl_v" field.
func DblVNotNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDblV)))
	})
}

// JSONVEQ applies the EQ predicate on the "json_v" field.
func JSONVEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJSONV), v))
	})
}

// JSONVNEQ applies the NEQ predicate on the "json_v" field.
func JSONVNEQ(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJSONV), v))
	})
}

// JSONVIn applies the In predicate on the "json_v" field.
func JSONVIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJSONV), v...))
	})
}

// JSONVNotIn applies the NotIn predicate on the "json_v" field.
func JSONVNotIn(vs ...string) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJSONV), v...))
	})
}

// JSONVGT applies the GT predicate on the "json_v" field.
func JSONVGT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJSONV), v))
	})
}

// JSONVGTE applies the GTE predicate on the "json_v" field.
func JSONVGTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJSONV), v))
	})
}

// JSONVLT applies the LT predicate on the "json_v" field.
func JSONVLT(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJSONV), v))
	})
}

// JSONVLTE applies the LTE predicate on the "json_v" field.
func JSONVLTE(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJSONV), v))
	})
}

// JSONVContains applies the Contains predicate on the "json_v" field.
func JSONVContains(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldJSONV), v))
	})
}

// JSONVHasPrefix applies the HasPrefix predicate on the "json_v" field.
func JSONVHasPrefix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldJSONV), v))
	})
}

// JSONVHasSuffix applies the HasSuffix predicate on the "json_v" field.
func JSONVHasSuffix(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldJSONV), v))
	})
}

// JSONVIsNil applies the IsNil predicate on the "json_v" field.
func JSONVIsNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldJSONV)))
	})
}

// JSONVNotNil applies the NotNil predicate on the "json_v" field.
func JSONVNotNil() predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldJSONV)))
	})
}

// JSONVEqualFold applies the EqualFold predicate on the "json_v" field.
func JSONVEqualFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldJSONV), v))
	})
}

// JSONVContainsFold applies the ContainsFold predicate on the "json_v" field.
func JSONVContainsFold(v string) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldJSONV), v))
	})
}

// LastUpdateTsEQ applies the EQ predicate on the "last_update_ts" field.
func LastUpdateTsEQ(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastUpdateTs), v))
	})
}

// LastUpdateTsNEQ applies the NEQ predicate on the "last_update_ts" field.
func LastUpdateTsNEQ(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastUpdateTs), v))
	})
}

// LastUpdateTsIn applies the In predicate on the "last_update_ts" field.
func LastUpdateTsIn(vs ...int64) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastUpdateTs), v...))
	})
}

// LastUpdateTsNotIn applies the NotIn predicate on the "last_update_ts" field.
func LastUpdateTsNotIn(vs ...int64) predicate.AttributeKV {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttributeKV(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastUpdateTs), v...))
	})
}

// LastUpdateTsGT applies the GT predicate on the "last_update_ts" field.
func LastUpdateTsGT(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastUpdateTs), v))
	})
}

// LastUpdateTsGTE applies the GTE predicate on the "last_update_ts" field.
func LastUpdateTsGTE(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastUpdateTs), v))
	})
}

// LastUpdateTsLT applies the LT predicate on the "last_update_ts" field.
func LastUpdateTsLT(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastUpdateTs), v))
	})
}

// LastUpdateTsLTE applies the LTE predicate on the "last_update_ts" field.
func LastUpdateTsLTE(v int64) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastUpdateTs), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AttributeKV) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AttributeKV) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AttributeKV) predicate.AttributeKV {
	return predicate.AttributeKV(func(s *sql.Selector) {
		p(s.Not())
	})
}
