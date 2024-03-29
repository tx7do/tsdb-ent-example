// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tsdb-ent-example/ent/predicate"
	"tsdb-ent-example/ent/sensordata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SensorDataUpdate is the builder for updating SensorData entities.
type SensorDataUpdate struct {
	config
	hooks    []Hook
	mutation *SensorDataMutation
}

// Where appends a list predicates to the SensorDataUpdate builder.
func (sdu *SensorDataUpdate) Where(ps ...predicate.SensorData) *SensorDataUpdate {
	sdu.mutation.Where(ps...)
	return sdu
}

// SetTime sets the "time" field.
func (sdu *SensorDataUpdate) SetTime(t time.Time) *SensorDataUpdate {
	sdu.mutation.SetTime(t)
	return sdu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (sdu *SensorDataUpdate) SetNillableTime(t *time.Time) *SensorDataUpdate {
	if t != nil {
		sdu.SetTime(*t)
	}
	return sdu
}

// ClearTime clears the value of the "time" field.
func (sdu *SensorDataUpdate) ClearTime() *SensorDataUpdate {
	sdu.mutation.ClearTime()
	return sdu
}

// SetSensorID sets the "sensor_id" field.
func (sdu *SensorDataUpdate) SetSensorID(i int) *SensorDataUpdate {
	sdu.mutation.ResetSensorID()
	sdu.mutation.SetSensorID(i)
	return sdu
}

// AddSensorID adds i to the "sensor_id" field.
func (sdu *SensorDataUpdate) AddSensorID(i int) *SensorDataUpdate {
	sdu.mutation.AddSensorID(i)
	return sdu
}

// SetTemperature sets the "temperature" field.
func (sdu *SensorDataUpdate) SetTemperature(f float64) *SensorDataUpdate {
	sdu.mutation.ResetTemperature()
	sdu.mutation.SetTemperature(f)
	return sdu
}

// AddTemperature adds f to the "temperature" field.
func (sdu *SensorDataUpdate) AddTemperature(f float64) *SensorDataUpdate {
	sdu.mutation.AddTemperature(f)
	return sdu
}

// SetCPU sets the "cpu" field.
func (sdu *SensorDataUpdate) SetCPU(f float64) *SensorDataUpdate {
	sdu.mutation.ResetCPU()
	sdu.mutation.SetCPU(f)
	return sdu
}

// AddCPU adds f to the "cpu" field.
func (sdu *SensorDataUpdate) AddCPU(f float64) *SensorDataUpdate {
	sdu.mutation.AddCPU(f)
	return sdu
}

// Mutation returns the SensorDataMutation object of the builder.
func (sdu *SensorDataUpdate) Mutation() *SensorDataMutation {
	return sdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdu *SensorDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdu.hooks) == 0 {
		affected, err = sdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdu.mutation = mutation
			affected, err = sdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdu.hooks) - 1; i >= 0; i-- {
			if sdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdu *SensorDataUpdate) SaveX(ctx context.Context) int {
	affected, err := sdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdu *SensorDataUpdate) Exec(ctx context.Context) error {
	_, err := sdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdu *SensorDataUpdate) ExecX(ctx context.Context) {
	if err := sdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sdu *SensorDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sensordata.Table,
			Columns: sensordata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensordata.FieldID,
			},
		},
	}
	if ps := sdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sensordata.FieldTime,
		})
	}
	if sdu.mutation.TimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sensordata.FieldTime,
		})
	}
	if value, ok := sdu.mutation.SensorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sensordata.FieldSensorID,
		})
	}
	if value, ok := sdu.mutation.AddedSensorID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sensordata.FieldSensorID,
		})
	}
	if value, ok := sdu.mutation.Temperature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldTemperature,
		})
	}
	if value, ok := sdu.mutation.AddedTemperature(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldTemperature,
		})
	}
	if value, ok := sdu.mutation.CPU(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldCPU,
		})
	}
	if value, ok := sdu.mutation.AddedCPU(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldCPU,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sensordata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SensorDataUpdateOne is the builder for updating a single SensorData entity.
type SensorDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SensorDataMutation
}

// SetTime sets the "time" field.
func (sduo *SensorDataUpdateOne) SetTime(t time.Time) *SensorDataUpdateOne {
	sduo.mutation.SetTime(t)
	return sduo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (sduo *SensorDataUpdateOne) SetNillableTime(t *time.Time) *SensorDataUpdateOne {
	if t != nil {
		sduo.SetTime(*t)
	}
	return sduo
}

// ClearTime clears the value of the "time" field.
func (sduo *SensorDataUpdateOne) ClearTime() *SensorDataUpdateOne {
	sduo.mutation.ClearTime()
	return sduo
}

// SetSensorID sets the "sensor_id" field.
func (sduo *SensorDataUpdateOne) SetSensorID(i int) *SensorDataUpdateOne {
	sduo.mutation.ResetSensorID()
	sduo.mutation.SetSensorID(i)
	return sduo
}

// AddSensorID adds i to the "sensor_id" field.
func (sduo *SensorDataUpdateOne) AddSensorID(i int) *SensorDataUpdateOne {
	sduo.mutation.AddSensorID(i)
	return sduo
}

// SetTemperature sets the "temperature" field.
func (sduo *SensorDataUpdateOne) SetTemperature(f float64) *SensorDataUpdateOne {
	sduo.mutation.ResetTemperature()
	sduo.mutation.SetTemperature(f)
	return sduo
}

// AddTemperature adds f to the "temperature" field.
func (sduo *SensorDataUpdateOne) AddTemperature(f float64) *SensorDataUpdateOne {
	sduo.mutation.AddTemperature(f)
	return sduo
}

// SetCPU sets the "cpu" field.
func (sduo *SensorDataUpdateOne) SetCPU(f float64) *SensorDataUpdateOne {
	sduo.mutation.ResetCPU()
	sduo.mutation.SetCPU(f)
	return sduo
}

// AddCPU adds f to the "cpu" field.
func (sduo *SensorDataUpdateOne) AddCPU(f float64) *SensorDataUpdateOne {
	sduo.mutation.AddCPU(f)
	return sduo
}

// Mutation returns the SensorDataMutation object of the builder.
func (sduo *SensorDataUpdateOne) Mutation() *SensorDataMutation {
	return sduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sduo *SensorDataUpdateOne) Select(field string, fields ...string) *SensorDataUpdateOne {
	sduo.fields = append([]string{field}, fields...)
	return sduo
}

// Save executes the query and returns the updated SensorData entity.
func (sduo *SensorDataUpdateOne) Save(ctx context.Context) (*SensorData, error) {
	var (
		err  error
		node *SensorData
	)
	if len(sduo.hooks) == 0 {
		node, err = sduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sduo.mutation = mutation
			node, err = sduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sduo.hooks) - 1; i >= 0; i-- {
			if sduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sduo *SensorDataUpdateOne) SaveX(ctx context.Context) *SensorData {
	node, err := sduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sduo *SensorDataUpdateOne) Exec(ctx context.Context) error {
	_, err := sduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sduo *SensorDataUpdateOne) ExecX(ctx context.Context) {
	if err := sduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sduo *SensorDataUpdateOne) sqlSave(ctx context.Context) (_node *SensorData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sensordata.Table,
			Columns: sensordata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensordata.FieldID,
			},
		},
	}
	id, ok := sduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SensorData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sensordata.FieldID)
		for _, f := range fields {
			if !sensordata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sensordata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sduo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sensordata.FieldTime,
		})
	}
	if sduo.mutation.TimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sensordata.FieldTime,
		})
	}
	if value, ok := sduo.mutation.SensorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sensordata.FieldSensorID,
		})
	}
	if value, ok := sduo.mutation.AddedSensorID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sensordata.FieldSensorID,
		})
	}
	if value, ok := sduo.mutation.Temperature(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldTemperature,
		})
	}
	if value, ok := sduo.mutation.AddedTemperature(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldTemperature,
		})
	}
	if value, ok := sduo.mutation.CPU(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldCPU,
		})
	}
	if value, ok := sduo.mutation.AddedCPU(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldCPU,
		})
	}
	_node = &SensorData{config: sduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sensordata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
