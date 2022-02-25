// Code generated by entc, DO NOT EDIT.

package ent

import (
	"tsdb-ent-example/ent/attributekv"
	"tsdb-ent-example/ent/sensor"
	"tsdb-ent-example/ent/sensordata"
	"tsdb-ent-example/ent/tskv"
	"tsdb-ent-example/ent/tskvdictionary"
	"tsdb-ent-example/ent/tskvlatest"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 6)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   attributekv.Table,
			Columns: attributekv.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: attributekv.FieldID,
			},
		},
		Type: "AttributeKV",
		Fields: map[string]*sqlgraph.FieldSpec{
			attributekv.FieldEntityType:    {Type: field.TypeString, Column: attributekv.FieldEntityType},
			attributekv.FieldEntityID:      {Type: field.TypeUUID, Column: attributekv.FieldEntityID},
			attributekv.FieldAttributeType: {Type: field.TypeString, Column: attributekv.FieldAttributeType},
			attributekv.FieldAttributeKey:  {Type: field.TypeString, Column: attributekv.FieldAttributeKey},
			attributekv.FieldBoolV:         {Type: field.TypeBool, Column: attributekv.FieldBoolV},
			attributekv.FieldStrV:          {Type: field.TypeString, Column: attributekv.FieldStrV},
			attributekv.FieldLongV:         {Type: field.TypeInt64, Column: attributekv.FieldLongV},
			attributekv.FieldDblV:          {Type: field.TypeFloat64, Column: attributekv.FieldDblV},
			attributekv.FieldJSONV:         {Type: field.TypeString, Column: attributekv.FieldJSONV},
			attributekv.FieldLastUpdateTs:  {Type: field.TypeInt64, Column: attributekv.FieldLastUpdateTs},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   sensor.Table,
			Columns: sensor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensor.FieldID,
			},
		},
		Type: "Sensor",
		Fields: map[string]*sqlgraph.FieldSpec{
			sensor.FieldType:     {Type: field.TypeString, Column: sensor.FieldType},
			sensor.FieldLocation: {Type: field.TypeString, Column: sensor.FieldLocation},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   sensordata.Table,
			Columns: sensordata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensordata.FieldID,
			},
		},
		Type: "SensorData",
		Fields: map[string]*sqlgraph.FieldSpec{
			sensordata.FieldTime:        {Type: field.TypeTime, Column: sensordata.FieldTime},
			sensordata.FieldSensorID:    {Type: field.TypeInt, Column: sensordata.FieldSensorID},
			sensordata.FieldTemperature: {Type: field.TypeFloat64, Column: sensordata.FieldTemperature},
			sensordata.FieldCPU:         {Type: field.TypeFloat64, Column: sensordata.FieldCPU},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tskv.Table,
			Columns: tskv.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tskv.FieldID,
			},
		},
		Type: "TsKv",
		Fields: map[string]*sqlgraph.FieldSpec{
			tskv.FieldEntityID: {Type: field.TypeUUID, Column: tskv.FieldEntityID},
			tskv.FieldKey:      {Type: field.TypeInt, Column: tskv.FieldKey},
			tskv.FieldBoolV:    {Type: field.TypeBool, Column: tskv.FieldBoolV},
			tskv.FieldStrV:     {Type: field.TypeString, Column: tskv.FieldStrV},
			tskv.FieldLongV:    {Type: field.TypeInt64, Column: tskv.FieldLongV},
			tskv.FieldDblV:     {Type: field.TypeFloat64, Column: tskv.FieldDblV},
			tskv.FieldJSONV:    {Type: field.TypeString, Column: tskv.FieldJSONV},
			tskv.FieldTs:       {Type: field.TypeInt64, Column: tskv.FieldTs},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tskvdictionary.Table,
			Columns: tskvdictionary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tskvdictionary.FieldID,
			},
		},
		Type: "TsKvDictionary",
		Fields: map[string]*sqlgraph.FieldSpec{
			tskvdictionary.FieldKey: {Type: field.TypeString, Column: tskvdictionary.FieldKey},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tskvlatest.Table,
			Columns: tskvlatest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tskvlatest.FieldID,
			},
		},
		Type: "TsKvLatest",
		Fields: map[string]*sqlgraph.FieldSpec{
			tskvlatest.FieldEntityID: {Type: field.TypeUUID, Column: tskvlatest.FieldEntityID},
			tskvlatest.FieldKey:      {Type: field.TypeInt, Column: tskvlatest.FieldKey},
			tskvlatest.FieldBoolV:    {Type: field.TypeBool, Column: tskvlatest.FieldBoolV},
			tskvlatest.FieldStrV:     {Type: field.TypeString, Column: tskvlatest.FieldStrV},
			tskvlatest.FieldLongV:    {Type: field.TypeInt64, Column: tskvlatest.FieldLongV},
			tskvlatest.FieldDblV:     {Type: field.TypeFloat64, Column: tskvlatest.FieldDblV},
			tskvlatest.FieldJSONV:    {Type: field.TypeString, Column: tskvlatest.FieldJSONV},
			tskvlatest.FieldTs:       {Type: field.TypeInt64, Column: tskvlatest.FieldTs},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (akq *AttributeKVQuery) addPredicate(pred func(s *sql.Selector)) {
	akq.predicates = append(akq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AttributeKVQuery builder.
func (akq *AttributeKVQuery) Filter() *AttributeKVFilter {
	return &AttributeKVFilter{akq}
}

// addPredicate implements the predicateAdder interface.
func (m *AttributeKVMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AttributeKVMutation builder.
func (m *AttributeKVMutation) Filter() *AttributeKVFilter {
	return &AttributeKVFilter{m}
}

// AttributeKVFilter provides a generic filtering capability at runtime for AttributeKVQuery.
type AttributeKVFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *AttributeKVFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *AttributeKVFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(attributekv.FieldID))
}

// WhereEntityType applies the entql string predicate on the entity_type field.
func (f *AttributeKVFilter) WhereEntityType(p entql.StringP) {
	f.Where(p.Field(attributekv.FieldEntityType))
}

// WhereEntityID applies the entql [16]byte predicate on the entity_id field.
func (f *AttributeKVFilter) WhereEntityID(p entql.ValueP) {
	f.Where(p.Field(attributekv.FieldEntityID))
}

// WhereAttributeType applies the entql string predicate on the attribute_type field.
func (f *AttributeKVFilter) WhereAttributeType(p entql.StringP) {
	f.Where(p.Field(attributekv.FieldAttributeType))
}

// WhereAttributeKey applies the entql string predicate on the attribute_key field.
func (f *AttributeKVFilter) WhereAttributeKey(p entql.StringP) {
	f.Where(p.Field(attributekv.FieldAttributeKey))
}

// WhereBoolV applies the entql bool predicate on the bool_v field.
func (f *AttributeKVFilter) WhereBoolV(p entql.BoolP) {
	f.Where(p.Field(attributekv.FieldBoolV))
}

// WhereStrV applies the entql string predicate on the str_v field.
func (f *AttributeKVFilter) WhereStrV(p entql.StringP) {
	f.Where(p.Field(attributekv.FieldStrV))
}

// WhereLongV applies the entql int64 predicate on the long_v field.
func (f *AttributeKVFilter) WhereLongV(p entql.Int64P) {
	f.Where(p.Field(attributekv.FieldLongV))
}

// WhereDblV applies the entql float64 predicate on the dbl_v field.
func (f *AttributeKVFilter) WhereDblV(p entql.Float64P) {
	f.Where(p.Field(attributekv.FieldDblV))
}

// WhereJSONV applies the entql string predicate on the json_v field.
func (f *AttributeKVFilter) WhereJSONV(p entql.StringP) {
	f.Where(p.Field(attributekv.FieldJSONV))
}

// WhereLastUpdateTs applies the entql int64 predicate on the last_update_ts field.
func (f *AttributeKVFilter) WhereLastUpdateTs(p entql.Int64P) {
	f.Where(p.Field(attributekv.FieldLastUpdateTs))
}

// addPredicate implements the predicateAdder interface.
func (sq *SensorQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SensorQuery builder.
func (sq *SensorQuery) Filter() *SensorFilter {
	return &SensorFilter{sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SensorMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SensorMutation builder.
func (m *SensorMutation) Filter() *SensorFilter {
	return &SensorFilter{m}
}

// SensorFilter provides a generic filtering capability at runtime for SensorQuery.
type SensorFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *SensorFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *SensorFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(sensor.FieldID))
}

// WhereType applies the entql string predicate on the type field.
func (f *SensorFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(sensor.FieldType))
}

// WhereLocation applies the entql string predicate on the location field.
func (f *SensorFilter) WhereLocation(p entql.StringP) {
	f.Where(p.Field(sensor.FieldLocation))
}

// addPredicate implements the predicateAdder interface.
func (sdq *SensorDataQuery) addPredicate(pred func(s *sql.Selector)) {
	sdq.predicates = append(sdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SensorDataQuery builder.
func (sdq *SensorDataQuery) Filter() *SensorDataFilter {
	return &SensorDataFilter{sdq}
}

// addPredicate implements the predicateAdder interface.
func (m *SensorDataMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SensorDataMutation builder.
func (m *SensorDataMutation) Filter() *SensorDataFilter {
	return &SensorDataFilter{m}
}

// SensorDataFilter provides a generic filtering capability at runtime for SensorDataQuery.
type SensorDataFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *SensorDataFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *SensorDataFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(sensordata.FieldID))
}

// WhereTime applies the entql time.Time predicate on the time field.
func (f *SensorDataFilter) WhereTime(p entql.TimeP) {
	f.Where(p.Field(sensordata.FieldTime))
}

// WhereSensorID applies the entql int predicate on the sensor_id field.
func (f *SensorDataFilter) WhereSensorID(p entql.IntP) {
	f.Where(p.Field(sensordata.FieldSensorID))
}

// WhereTemperature applies the entql float64 predicate on the temperature field.
func (f *SensorDataFilter) WhereTemperature(p entql.Float64P) {
	f.Where(p.Field(sensordata.FieldTemperature))
}

// WhereCPU applies the entql float64 predicate on the cpu field.
func (f *SensorDataFilter) WhereCPU(p entql.Float64P) {
	f.Where(p.Field(sensordata.FieldCPU))
}

// addPredicate implements the predicateAdder interface.
func (tkq *TsKvQuery) addPredicate(pred func(s *sql.Selector)) {
	tkq.predicates = append(tkq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TsKvQuery builder.
func (tkq *TsKvQuery) Filter() *TsKvFilter {
	return &TsKvFilter{tkq}
}

// addPredicate implements the predicateAdder interface.
func (m *TsKvMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TsKvMutation builder.
func (m *TsKvMutation) Filter() *TsKvFilter {
	return &TsKvFilter{m}
}

// TsKvFilter provides a generic filtering capability at runtime for TsKvQuery.
type TsKvFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *TsKvFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TsKvFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tskv.FieldID))
}

// WhereEntityID applies the entql [16]byte predicate on the entity_id field.
func (f *TsKvFilter) WhereEntityID(p entql.ValueP) {
	f.Where(p.Field(tskv.FieldEntityID))
}

// WhereKey applies the entql int predicate on the key field.
func (f *TsKvFilter) WhereKey(p entql.IntP) {
	f.Where(p.Field(tskv.FieldKey))
}

// WhereBoolV applies the entql bool predicate on the bool_v field.
func (f *TsKvFilter) WhereBoolV(p entql.BoolP) {
	f.Where(p.Field(tskv.FieldBoolV))
}

// WhereStrV applies the entql string predicate on the str_v field.
func (f *TsKvFilter) WhereStrV(p entql.StringP) {
	f.Where(p.Field(tskv.FieldStrV))
}

// WhereLongV applies the entql int64 predicate on the long_v field.
func (f *TsKvFilter) WhereLongV(p entql.Int64P) {
	f.Where(p.Field(tskv.FieldLongV))
}

// WhereDblV applies the entql float64 predicate on the dbl_v field.
func (f *TsKvFilter) WhereDblV(p entql.Float64P) {
	f.Where(p.Field(tskv.FieldDblV))
}

// WhereJSONV applies the entql string predicate on the json_v field.
func (f *TsKvFilter) WhereJSONV(p entql.StringP) {
	f.Where(p.Field(tskv.FieldJSONV))
}

// WhereTs applies the entql int64 predicate on the ts field.
func (f *TsKvFilter) WhereTs(p entql.Int64P) {
	f.Where(p.Field(tskv.FieldTs))
}

// addPredicate implements the predicateAdder interface.
func (tkdq *TsKvDictionaryQuery) addPredicate(pred func(s *sql.Selector)) {
	tkdq.predicates = append(tkdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TsKvDictionaryQuery builder.
func (tkdq *TsKvDictionaryQuery) Filter() *TsKvDictionaryFilter {
	return &TsKvDictionaryFilter{tkdq}
}

// addPredicate implements the predicateAdder interface.
func (m *TsKvDictionaryMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TsKvDictionaryMutation builder.
func (m *TsKvDictionaryMutation) Filter() *TsKvDictionaryFilter {
	return &TsKvDictionaryFilter{m}
}

// TsKvDictionaryFilter provides a generic filtering capability at runtime for TsKvDictionaryQuery.
type TsKvDictionaryFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *TsKvDictionaryFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TsKvDictionaryFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tskvdictionary.FieldID))
}

// WhereKey applies the entql string predicate on the key field.
func (f *TsKvDictionaryFilter) WhereKey(p entql.StringP) {
	f.Where(p.Field(tskvdictionary.FieldKey))
}

// addPredicate implements the predicateAdder interface.
func (tklq *TsKvLatestQuery) addPredicate(pred func(s *sql.Selector)) {
	tklq.predicates = append(tklq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TsKvLatestQuery builder.
func (tklq *TsKvLatestQuery) Filter() *TsKvLatestFilter {
	return &TsKvLatestFilter{tklq}
}

// addPredicate implements the predicateAdder interface.
func (m *TsKvLatestMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TsKvLatestMutation builder.
func (m *TsKvLatestMutation) Filter() *TsKvLatestFilter {
	return &TsKvLatestFilter{m}
}

// TsKvLatestFilter provides a generic filtering capability at runtime for TsKvLatestQuery.
type TsKvLatestFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *TsKvLatestFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TsKvLatestFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tskvlatest.FieldID))
}

// WhereEntityID applies the entql [16]byte predicate on the entity_id field.
func (f *TsKvLatestFilter) WhereEntityID(p entql.ValueP) {
	f.Where(p.Field(tskvlatest.FieldEntityID))
}

// WhereKey applies the entql int predicate on the key field.
func (f *TsKvLatestFilter) WhereKey(p entql.IntP) {
	f.Where(p.Field(tskvlatest.FieldKey))
}

// WhereBoolV applies the entql bool predicate on the bool_v field.
func (f *TsKvLatestFilter) WhereBoolV(p entql.BoolP) {
	f.Where(p.Field(tskvlatest.FieldBoolV))
}

// WhereStrV applies the entql string predicate on the str_v field.
func (f *TsKvLatestFilter) WhereStrV(p entql.StringP) {
	f.Where(p.Field(tskvlatest.FieldStrV))
}

// WhereLongV applies the entql int64 predicate on the long_v field.
func (f *TsKvLatestFilter) WhereLongV(p entql.Int64P) {
	f.Where(p.Field(tskvlatest.FieldLongV))
}

// WhereDblV applies the entql float64 predicate on the dbl_v field.
func (f *TsKvLatestFilter) WhereDblV(p entql.Float64P) {
	f.Where(p.Field(tskvlatest.FieldDblV))
}

// WhereJSONV applies the entql string predicate on the json_v field.
func (f *TsKvLatestFilter) WhereJSONV(p entql.StringP) {
	f.Where(p.Field(tskvlatest.FieldJSONV))
}

// WhereTs applies the entql int64 predicate on the ts field.
func (f *TsKvLatestFilter) WhereTs(p entql.Int64P) {
	f.Where(p.Field(tskvlatest.FieldTs))
}