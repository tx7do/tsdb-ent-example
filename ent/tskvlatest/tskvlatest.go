// Code generated by entc, DO NOT EDIT.

package tskvlatest

const (
	// Label holds the string label denoting the tskvlatest type in the database.
	Label = "ts_kv_latest"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
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
	// FieldTs holds the string denoting the ts field in the database.
	FieldTs = "ts"
	// Table holds the table name of the tskvlatest in the database.
	Table = "ts_kv_latest"
)

// Columns holds all SQL columns for tskvlatest fields.
var Columns = []string{
	FieldID,
	FieldEntityID,
	FieldKey,
	FieldBoolV,
	FieldStrV,
	FieldLongV,
	FieldDblV,
	FieldJSONV,
	FieldTs,
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
	// StrVValidator is a validator for the "str_v" field. It is called by the builders before save.
	StrVValidator func(string) error
)
