package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// AttributeKV holds the schema definition for the AttributeKV entity.
type AttributeKV struct {
	ent.Schema
}

func (AttributeKV) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "attribute_kv",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the AttributeKV.
func (AttributeKV) Fields() []ent.Field {
	return []ent.Field{
		field.String("entity_type").
			Comment("实体类型").
			MaxLen(255),
		field.UUID("entity_id", uuid.UUID{}).
			Comment("实体ID"),
		field.String("attribute_type").
			Comment("属性类型").
			MaxLen(255),
		field.String("attribute_key").
			Comment("属性键名").
			MaxLen(255),
		field.Bool("bool_v").
			Comment("BOOL值").
			Nillable().
			Optional(),
		field.String("str_v").
			Comment("STRING值").
			MaxLen(10000000).
			Nillable().
			Optional(),
		field.Int64("long_v").
			Comment("LONG值").
			Nillable().
			Optional(),
		field.Float("dbl_v").
			Comment("DOUBLE值").
			Nillable().
			Optional(),
		field.String("json_v").
			Comment("JSON值").
			Nillable().
			Optional(),
		field.Int64("last_update_ts").
			Comment("最后更新的时间戳"),
	}
}

// Edges of the AttributeKV.
func (AttributeKV) Edges() []ent.Edge {
	return nil
}

// Indexes of the AttributeKV.
func (AttributeKV) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entity_type", "entity_id", "attribute_type", "attribute_key").
			StorageKey("attribute_kv_pkey"),
		index.Fields("entity_id", "attribute_key", "last_update_ts").
			StorageKey("idx_attribute_kv_by_key_and_last_update_ts"),
	}
}
