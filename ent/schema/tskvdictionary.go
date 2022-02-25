package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TsKvDictionary holds the schema definition for the TsKvDictionary entity.
type TsKvDictionary struct {
	ent.Schema
}

func (TsKvDictionary) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "ts_kv_dictionary",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
	}
}

// Fields of the TsKvDictionary.
func (TsKvDictionary) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Comment("键名").
			MaxLen(255).NotEmpty(),
	}
}

// Edges of the TsKvDictionary.
func (TsKvDictionary) Edges() []ent.Edge {
	return nil
}

// Indexes of the TsKvDictionary.
func (TsKvDictionary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").StorageKey("ts_key_id_pkey"),
	}
}
