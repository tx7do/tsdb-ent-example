package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// TsKvLatest holds the schema definition for the TsKvLatest entity.
type TsKvLatest struct {
	ent.Schema
}

func (TsKvLatest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "ts_kv_latest",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the TsKvLatest.
func (TsKvLatest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("entity_id", uuid.UUID{}).
			Comment("实体ID"),
		field.Int("key").
			Comment("键ID"),
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
		field.Int64("ts").
			Comment("最后更新的时间戳"),
	}
}

// Edges of the TsKvLatest.
func (TsKvLatest) Edges() []ent.Edge {
	return nil
}

// Indexes of the TsKvLatest.
func (TsKvLatest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("entity_id", "key").StorageKey("ts_kv_latest_pkey"),
	}
}
