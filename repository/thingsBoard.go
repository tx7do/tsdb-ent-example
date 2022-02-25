package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/lib/pq"
	"tsdb-ent-example/ent"
	"tsdb-ent-example/ent/tskvdictionary"
	"tsdb-ent-example/model"
)

type ThingsBoardRepository struct {
	client *ent.Client
	ctx    context.Context
}

func NewThingsBoardRepository() *ThingsBoardRepository {
	r := &ThingsBoardRepository{}
	_ = r.Connect()
	return r
}

func (r *ThingsBoardRepository) Connect() error {
	cli, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=thingsboard sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Run the auto migration tool.
	//if err := cli.Schema.Create(context.Background()); err != nil {
	//	fmt.Printf("failed creating schema resources: %v \n", err)
	//}

	r.ctx = context.Background()
	r.client = cli
	return nil
}

func (r *ThingsBoardRepository) GetOrSaveKeyId(key string) int {
	keyId := r.GetKeyId(key)
	if keyId == -1 {
		if r.SaveKeyId(key) == nil {
			return r.GetKeyId(key)
		} else {
			return -1
		}
	} else {
		return keyId
	}
}

func (r *ThingsBoardRepository) GetKeyId(key string) int {
	KeyId, err := r.client.TsKvDictionary.Query().
		Where(tskvdictionary.KeyEQ(key)).
		Select(tskvdictionary.FieldID).
		Int(r.ctx)
	if err != nil {
		return -1
	} else {
		return KeyId
	}
}

func (r *ThingsBoardRepository) SaveKeyId(key string) error {
	return r.client.TsKvDictionary.Create().
		SetKey(key).
		Exec(r.ctx)
}

func (r *ThingsBoardRepository) SaveOrUpdateTsKv(value *model.TsKv) error {
	err := r.client.Debug().TsKv.Create().
		SetEntityID(value.EntityId).
		SetKey(value.Key).
		SetTs(value.Timestamp).
		SetNillableBoolV(value.BoolV).
		SetNillableStrV(value.StringV).
		SetNillableLongV(value.LongV).
		SetNillableDblV(value.DoubleV).
		SetNillableJSONV(value.JsonV).
		OnConflict(sql.ConflictColumns("entity_id", "key", "ts")).
		Update(func(u *ent.TsKvUpsert) {
			if value.BoolV != nil {
				u.SetBoolV(*value.BoolV)
			} else {
				u.ClearBoolV()
			}
			if value.StringV != nil {
				u.SetStrV(*value.StringV)
			} else {
				u.ClearStrV()
			}
			if value.LongV != nil {
				u.SetLongV(*value.LongV)
			} else {
				u.ClearLongV()
			}
			if value.DoubleV != nil {
				u.SetDblV(*value.DoubleV)
			} else {
				u.ClearDblV()
			}
			if value.JsonV != nil {
				u.SetJSONV(*value.JsonV)
			} else {
				u.ClearJSONV()
			}
		}).
		Exec(r.ctx)
	return err
}

func (r *ThingsBoardRepository) SaveOrUpdateTsKvLatest(value *model.TsKv) error {
	err := r.client.Debug().TsKvLatest.Create().
		SetEntityID(value.EntityId).
		SetKey(value.Key).
		SetTs(value.Timestamp).
		SetNillableBoolV(value.BoolV).
		SetNillableStrV(value.StringV).
		SetNillableLongV(value.LongV).
		SetNillableDblV(value.DoubleV).
		SetNillableJSONV(value.JsonV).
		OnConflict(sql.ConflictColumns("entity_id", "key")).
		Update(func(u *ent.TsKvLatestUpsert) {
			u.
				SetTs(value.Timestamp)

			if value.BoolV != nil {
				u.SetBoolV(*value.BoolV)
			} else {
				u.ClearBoolV()
			}
			if value.StringV != nil {
				u.SetStrV(*value.StringV)
			} else {
				u.ClearStrV()
			}
			if value.LongV != nil {
				u.SetLongV(*value.LongV)
			} else {
				u.ClearLongV()
			}
			if value.DoubleV != nil {
				u.SetDblV(*value.DoubleV)
			} else {
				u.ClearDblV()
			}
			if value.JsonV != nil {
				u.SetJSONV(*value.JsonV)
			} else {
				u.ClearJSONV()
			}
		}).
		Exec(r.ctx)
	return err
}
