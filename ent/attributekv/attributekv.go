// Code generated by entc, DO NOT EDIT.

package attributekv

const (
	// Label holds the string label denoting the attributekv type in the database.
	Label = "attribute_kv"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntityType holds the string denoting the entity_type field in the database.
	FieldEntityType = "entity_type"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// FieldAttributeType holds the string denoting the attribute_type field in the database.
	FieldAttributeType = "attribute_type"
	// FieldAttributeKey holds the string denoting the attribute_key field in the database.
	FieldAttributeKey = "attribute_key"
	// FieldBoolV holds the string denoting the bool_v field in the database.
	FieldBoolV = "bool_v"
	// FieldStrV holds the string denoting the str_v field in the database.
	FieldStrV = "str_v"
	// FieldLongV holds the string denoting the long_v field in the database.
	FieldLongV = "long_v"
	// FieldDblV holds the string denoting the dbl_v field in the database.
	FieldDblV = "dbl_v"
	// FieldJSONV holds the string denoting the json_v field in the database.
	FieldJSONV = "json_v"
	// FieldLastUpdateTs holds the string denoting the last_update_ts field in the database.
	FieldLastUpdateTs = "last_update_ts"
	// Table holds the table name of the attributekv in the database.
	Table = "attribute_kv"
)

// Columns holds all SQL columns for attributekv fields.
var Columns = []string{
	FieldID,
	FieldEntityType,
	FieldEntityID,
	FieldAttributeType,
	FieldAttributeKey,
	FieldBoolV,
	FieldStrV,
	FieldLongV,
	FieldDblV,
	FieldJSONV,
	FieldLastUpdateTs,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EntityTypeValidator is a validator for the "entity_type" field. It is called by the builders before save.
	EntityTypeValidator func(string) error
	// AttributeTypeValidator is a validator for the "attribute_type" field. It is called by the builders before save.
	AttributeTypeValidator func(string) error
	// AttributeKeyValidator is a validator for the "attribute_key" field. It is called by the builders before save.
	AttributeKeyValidator func(string) error
	// StrVValidator is a validator for the "str_v" field. It is called by the builders before save.
	StrVValidator func(string) error
)