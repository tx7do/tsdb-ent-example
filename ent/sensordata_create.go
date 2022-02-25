// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tsdb-ent-example/ent/sensordata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SensorDataCreate is the builder for creating a SensorData entity.
type SensorDataCreate struct {
	config
	mutation *SensorDataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTime sets the "time" field.
func (sdc *SensorDataCreate) SetTime(t time.Time) *SensorDataCreate {
	sdc.mutation.SetTime(t)
	return sdc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (sdc *SensorDataCreate) SetNillableTime(t *time.Time) *SensorDataCreate {
	if t != nil {
		sdc.SetTime(*t)
	}
	return sdc
}

// SetSensorID sets the "sensor_id" field.
func (sdc *SensorDataCreate) SetSensorID(i int) *SensorDataCreate {
	sdc.mutation.SetSensorID(i)
	return sdc
}

// SetTemperature sets the "temperature" field.
func (sdc *SensorDataCreate) SetTemperature(f float64) *SensorDataCreate {
	sdc.mutation.SetTemperature(f)
	return sdc
}

// SetCPU sets the "cpu" field.
func (sdc *SensorDataCreate) SetCPU(f float64) *SensorDataCreate {
	sdc.mutation.SetCPU(f)
	return sdc
}

// Mutation returns the SensorDataMutation object of the builder.
func (sdc *SensorDataCreate) Mutation() *SensorDataMutation {
	return sdc.mutation
}

// Save creates the SensorData in the database.
func (sdc *SensorDataCreate) Save(ctx context.Context) (*SensorData, error) {
	var (
		err  error
		node *SensorData
	)
	if len(sdc.hooks) == 0 {
		if err = sdc.check(); err != nil {
			return nil, err
		}
		node, err = sdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdc.check(); err != nil {
				return nil, err
			}
			sdc.mutation = mutation
			if node, err = sdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sdc.hooks) - 1; i >= 0; i-- {
			if sdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SensorDataCreate) SaveX(ctx context.Context) *SensorData {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SensorDataCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SensorDataCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SensorDataCreate) check() error {
	if _, ok := sdc.mutation.SensorID(); !ok {
		return &ValidationError{Name: "sensor_id", err: errors.New(`ent: missing required field "SensorData.sensor_id"`)}
	}
	if _, ok := sdc.mutation.Temperature(); !ok {
		return &ValidationError{Name: "temperature", err: errors.New(`ent: missing required field "SensorData.temperature"`)}
	}
	if _, ok := sdc.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "SensorData.cpu"`)}
	}
	return nil
}

func (sdc *SensorDataCreate) sqlSave(ctx context.Context) (*SensorData, error) {
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sdc *SensorDataCreate) createSpec() (*SensorData, *sqlgraph.CreateSpec) {
	var (
		_node = &SensorData{config: sdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sensordata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensordata.FieldID,
			},
		}
	)
	_spec.OnConflict = sdc.conflict
	if value, ok := sdc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sensordata.FieldTime,
		})
		_node.Time = &value
	}
	if value, ok := sdc.mutation.SensorID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sensordata.FieldSensorID,
		})
		_node.SensorID = value
	}
	if value, ok := sdc.mutation.Temperature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldTemperature,
		})
		_node.Temperature = value
	}
	if value, ok := sdc.mutation.CPU(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sensordata.FieldCPU,
		})
		_node.CPU = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SensorData.Create().
//		SetTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SensorDataUpsert) {
//			SetTime(v+v).
//		}).
//		Exec(ctx)
//
func (sdc *SensorDataCreate) OnConflict(opts ...sql.ConflictOption) *SensorDataUpsertOne {
	sdc.conflict = opts
	return &SensorDataUpsertOne{
		create: sdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SensorData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sdc *SensorDataCreate) OnConflictColumns(columns ...string) *SensorDataUpsertOne {
	sdc.conflict = append(sdc.conflict, sql.ConflictColumns(columns...))
	return &SensorDataUpsertOne{
		create: sdc,
	}
}

type (
	// SensorDataUpsertOne is the builder for "upsert"-ing
	//  one SensorData node.
	SensorDataUpsertOne struct {
		create *SensorDataCreate
	}

	// SensorDataUpsert is the "OnConflict" setter.
	SensorDataUpsert struct {
		*sql.UpdateSet
	}
)

// SetTime sets the "time" field.
func (u *SensorDataUpsert) SetTime(v time.Time) *SensorDataUpsert {
	u.Set(sensordata.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *SensorDataUpsert) UpdateTime() *SensorDataUpsert {
	u.SetExcluded(sensordata.FieldTime)
	return u
}

// ClearTime clears the value of the "time" field.
func (u *SensorDataUpsert) ClearTime() *SensorDataUpsert {
	u.SetNull(sensordata.FieldTime)
	return u
}

// SetSensorID sets the "sensor_id" field.
func (u *SensorDataUpsert) SetSensorID(v int) *SensorDataUpsert {
	u.Set(sensordata.FieldSensorID, v)
	return u
}

// UpdateSensorID sets the "sensor_id" field to the value that was provided on create.
func (u *SensorDataUpsert) UpdateSensorID() *SensorDataUpsert {
	u.SetExcluded(sensordata.FieldSensorID)
	return u
}

// AddSensorID adds v to the "sensor_id" field.
func (u *SensorDataUpsert) AddSensorID(v int) *SensorDataUpsert {
	u.Add(sensordata.FieldSensorID, v)
	return u
}

// SetTemperature sets the "temperature" field.
func (u *SensorDataUpsert) SetTemperature(v float64) *SensorDataUpsert {
	u.Set(sensordata.FieldTemperature, v)
	return u
}

// UpdateTemperature sets the "temperature" field to the value that was provided on create.
func (u *SensorDataUpsert) UpdateTemperature() *SensorDataUpsert {
	u.SetExcluded(sensordata.FieldTemperature)
	return u
}

// AddTemperature adds v to the "temperature" field.
func (u *SensorDataUpsert) AddTemperature(v float64) *SensorDataUpsert {
	u.Add(sensordata.FieldTemperature, v)
	return u
}

// SetCPU sets the "cpu" field.
func (u *SensorDataUpsert) SetCPU(v float64) *SensorDataUpsert {
	u.Set(sensordata.FieldCPU, v)
	return u
}

// UpdateCPU sets the "cpu" field to the value that was provided on create.
func (u *SensorDataUpsert) UpdateCPU() *SensorDataUpsert {
	u.SetExcluded(sensordata.FieldCPU)
	return u
}

// AddCPU adds v to the "cpu" field.
func (u *SensorDataUpsert) AddCPU(v float64) *SensorDataUpsert {
	u.Add(sensordata.FieldCPU, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.SensorData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *SensorDataUpsertOne) UpdateNewValues() *SensorDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.SensorData.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *SensorDataUpsertOne) Ignore() *SensorDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SensorDataUpsertOne) DoNothing() *SensorDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SensorDataCreate.OnConflict
// documentation for more info.
func (u *SensorDataUpsertOne) Update(set func(*SensorDataUpsert)) *SensorDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SensorDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetTime sets the "time" field.
func (u *SensorDataUpsertOne) SetTime(v time.Time) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *SensorDataUpsertOne) UpdateTime() *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateTime()
	})
}

// ClearTime clears the value of the "time" field.
func (u *SensorDataUpsertOne) ClearTime() *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.ClearTime()
	})
}

// SetSensorID sets the "sensor_id" field.
func (u *SensorDataUpsertOne) SetSensorID(v int) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetSensorID(v)
	})
}

// AddSensorID adds v to the "sensor_id" field.
func (u *SensorDataUpsertOne) AddSensorID(v int) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.AddSensorID(v)
	})
}

// UpdateSensorID sets the "sensor_id" field to the value that was provided on create.
func (u *SensorDataUpsertOne) UpdateSensorID() *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateSensorID()
	})
}

// SetTemperature sets the "temperature" field.
func (u *SensorDataUpsertOne) SetTemperature(v float64) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetTemperature(v)
	})
}

// AddTemperature adds v to the "temperature" field.
func (u *SensorDataUpsertOne) AddTemperature(v float64) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.AddTemperature(v)
	})
}

// UpdateTemperature sets the "temperature" field to the value that was provided on create.
func (u *SensorDataUpsertOne) UpdateTemperature() *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateTemperature()
	})
}

// SetCPU sets the "cpu" field.
func (u *SensorDataUpsertOne) SetCPU(v float64) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetCPU(v)
	})
}

// AddCPU adds v to the "cpu" field.
func (u *SensorDataUpsertOne) AddCPU(v float64) *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.AddCPU(v)
	})
}

// UpdateCPU sets the "cpu" field to the value that was provided on create.
func (u *SensorDataUpsertOne) UpdateCPU() *SensorDataUpsertOne {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateCPU()
	})
}

// Exec executes the query.
func (u *SensorDataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SensorDataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SensorDataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SensorDataUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SensorDataUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SensorDataCreateBulk is the builder for creating many SensorData entities in bulk.
type SensorDataCreateBulk struct {
	config
	builders []*SensorDataCreate
	conflict []sql.ConflictOption
}

// Save creates the SensorData entities in the database.
func (sdcb *SensorDataCreateBulk) Save(ctx context.Context) ([]*SensorData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SensorData, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SensorDataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SensorDataCreateBulk) SaveX(ctx context.Context) []*SensorData {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SensorDataCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SensorDataCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SensorData.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SensorDataUpsert) {
//			SetTime(v+v).
//		}).
//		Exec(ctx)
//
func (sdcb *SensorDataCreateBulk) OnConflict(opts ...sql.ConflictOption) *SensorDataUpsertBulk {
	sdcb.conflict = opts
	return &SensorDataUpsertBulk{
		create: sdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SensorData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sdcb *SensorDataCreateBulk) OnConflictColumns(columns ...string) *SensorDataUpsertBulk {
	sdcb.conflict = append(sdcb.conflict, sql.ConflictColumns(columns...))
	return &SensorDataUpsertBulk{
		create: sdcb,
	}
}

// SensorDataUpsertBulk is the builder for "upsert"-ing
// a bulk of SensorData nodes.
type SensorDataUpsertBulk struct {
	create *SensorDataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SensorData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *SensorDataUpsertBulk) UpdateNewValues() *SensorDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SensorData.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *SensorDataUpsertBulk) Ignore() *SensorDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SensorDataUpsertBulk) DoNothing() *SensorDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SensorDataCreateBulk.OnConflict
// documentation for more info.
func (u *SensorDataUpsertBulk) Update(set func(*SensorDataUpsert)) *SensorDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SensorDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetTime sets the "time" field.
func (u *SensorDataUpsertBulk) SetTime(v time.Time) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *SensorDataUpsertBulk) UpdateTime() *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateTime()
	})
}

// ClearTime clears the value of the "time" field.
func (u *SensorDataUpsertBulk) ClearTime() *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.ClearTime()
	})
}

// SetSensorID sets the "sensor_id" field.
func (u *SensorDataUpsertBulk) SetSensorID(v int) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetSensorID(v)
	})
}

// AddSensorID adds v to the "sensor_id" field.
func (u *SensorDataUpsertBulk) AddSensorID(v int) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.AddSensorID(v)
	})
}

// UpdateSensorID sets the "sensor_id" field to the value that was provided on create.
func (u *SensorDataUpsertBulk) UpdateSensorID() *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateSensorID()
	})
}

// SetTemperature sets the "temperature" field.
func (u *SensorDataUpsertBulk) SetTemperature(v float64) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetTemperature(v)
	})
}

// AddTemperature adds v to the "temperature" field.
func (u *SensorDataUpsertBulk) AddTemperature(v float64) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.AddTemperature(v)
	})
}

// UpdateTemperature sets the "temperature" field to the value that was provided on create.
func (u *SensorDataUpsertBulk) UpdateTemperature() *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateTemperature()
	})
}

// SetCPU sets the "cpu" field.
func (u *SensorDataUpsertBulk) SetCPU(v float64) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.SetCPU(v)
	})
}

// AddCPU adds v to the "cpu" field.
func (u *SensorDataUpsertBulk) AddCPU(v float64) *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.AddCPU(v)
	})
}

// UpdateCPU sets the "cpu" field to the value that was provided on create.
func (u *SensorDataUpsertBulk) UpdateCPU() *SensorDataUpsertBulk {
	return u.Update(func(s *SensorDataUpsert) {
		s.UpdateCPU()
	})
}

// Exec executes the query.
func (u *SensorDataUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SensorDataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SensorDataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SensorDataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
