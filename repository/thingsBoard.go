package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"strconv"
	"tsdb-ent-example/ent"
	"tsdb-ent-example/ent/tskv"
	"tsdb-ent-example/ent/tskvdictionary"
	"tsdb-ent-example/ent/tskvlatest"
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

// GetOrSaveKeyId 查询和更新一条数据到 数据键名-键ID映射表
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

// GetKeyId 查询一条数据到 数据键名-键ID映射表
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

// SaveKeyId 插入一条数据到 数据键名-键ID映射表
func (r *ThingsBoardRepository) SaveKeyId(key string) error {
	return r.client.TsKvDictionary.Create().
		SetKey(key).
		Exec(r.ctx)
}

// SaveOrUpdateTsKv 插入一条数据到遥测数据历史数据表去
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
		OnConflict(sql.ConflictColumns(tskv.FieldEntityID, tskv.FieldKey, tskv.FieldTs)).
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

// SaveOrUpdateTsKvLatest 插入一条数据到最新遥测数据表去
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
		OnConflict(sql.ConflictColumns(tskvlatest.FieldEntityID, tskvlatest.FieldKey)).
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

// FindTsKvLatest 查询指定设备的最新遥测数据
func (r *ThingsBoardRepository) FindTsKvLatest(entityId uuid.UUID) ([]model.TsKvLatest, error) {
	var results []model.TsKvLatest
	err := r.client.Debug().TsKvLatest.Query().
		Where(
			func(s *sql.Selector) {
				t := sql.Table(tskvdictionary.Table)
				s.Join(t).On(s.C(tskvlatest.FieldKey), t.C(tskvdictionary.FieldID))
				s.Select(
					sql.As(t.C(tskvdictionary.FieldKey), "key"),

					sql.As(s.C(tskvlatest.FieldEntityID), "entity_id"),
					sql.As(s.C(tskvlatest.FieldKey), "key_id"),
					sql.As(s.C(tskvlatest.FieldStrV), "str_v"),
					sql.As(s.C(tskvlatest.FieldBoolV), "bool_v"),
					sql.As(s.C(tskvlatest.FieldLongV), "long_v"),
					sql.As(s.C(tskvlatest.FieldDblV), "dbl_v"),
					sql.As(s.C(tskvlatest.FieldJSONV), "json_v"),
					sql.As(s.C(tskvlatest.FieldTs), "ts"),
				)
			},
			tskvlatest.EntityIDEQ(entityId)).
		Select().
		Scan(r.ctx, &results)
	return results, err
}

// FindAllKeysByEntityIds 查询指定设备的所有遥测数据键名
func (r *ThingsBoardRepository) FindAllKeysByEntityIds(entityId []interface{}) ([]string, error) {
	keys, err := r.client.Debug().TsKvLatest.Query().
		Where(
			func(s *sql.Selector) {
				t := sql.Table(tskvdictionary.Table)

				s.Join(t).On(s.C(tskvlatest.FieldKey), t.C(tskvdictionary.FieldID))

				s.Select(
					sql.As(t.C(tskvdictionary.FieldKey), "strKey"),
				).
					Distinct().
					Where(
						sql.In(
							s.C(tskvlatest.FieldEntityID),
							entityId...,
						),
					).OrderBy(t.C(tskvdictionary.FieldKey))
			},
		).
		Select().
		Strings(r.ctx)
	return keys, err
}

type AggregationResult struct {
	TsBucket    int64   `json:"ts_bucket"`
	LongValue   int64   `json:"long_value"`
	DoubleValue float64 `json:"double_value"`
	StringValue string  `json:"string_value"`
	JsonValue   string  `json:"json_value"`

	BoolValueCount   int `json:"bool_value_count"`
	LongValueCount   int `json:"long_value_count"`
	DoubleValueCount int `json:"double_value_count"`
	StringValueCount int `json:"string_value_count"`
	JsonValueCount   int `json:"json_value_count"`
}

func (r *ThingsBoardRepository) FindAvg(entityId uuid.UUID, entityKeyId int, timeBucket int64, startTs, endTs int64) ([]AggregationResult, error) {
	strTimeBucket := strconv.FormatInt(timeBucket, 10)
	strStartTs := strconv.FormatInt(startTs, 10)

	builder := &sql.Builder{}

	var results []AggregationResult

	err := r.client.Debug().TsKv.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(
					builder.Reset().WriteString("time_bucket(").
						WriteString(strTimeBucket).Comma().
						WriteString(s.C(tskv.FieldTs)).Comma().
						WriteString(strStartTs).
						WriteByte(')').
						String(),
					"ts_bucket",
				),
				//sql.As(
				//	builder.Reset().WriteString(strTimeBucket).String(),
				//	"interval",
				//),
				sql.As(sql.Sum(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldLongV).Comma().
						WriteString("0").
						WriteByte(')').
						String(),
				),
					"long_value",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldDblV).Comma().
						WriteString("0.0").
						WriteByte(')').
						String(),
				),
					"double_value",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldLongV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"long_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldDblV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"double_value_Count",
				),
				//sql.As(`NULL`, "strValue"),
				//sql.As("AVG", "aggType"),
			).
				Where(
					sql.And(
						sql.EQ(tskv.FieldEntityID, entityId),
						sql.EQ(tskv.FieldKey, entityKeyId),
						sql.GT(tskv.FieldTs, startTs),
						sql.LTE(tskv.FieldTs, endTs),
					),
				).
				GroupBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket").
				OrderBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket")
		}).
		Scan(r.ctx, &results)
	return results, err
}

func (r *ThingsBoardRepository) FindMax(entityId uuid.UUID, entityKeyId int, timeBucket int64, startTs, endTs int64) ([]AggregationResult, error) {
	strTimeBucket := strconv.FormatInt(timeBucket, 10)
	strStartTs := strconv.FormatInt(startTs, 10)

	builder := &sql.Builder{}

	var results []AggregationResult

	err := r.client.Debug().TsKv.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(
					builder.Reset().WriteString("time_bucket(").
						WriteString(strTimeBucket).Comma().
						WriteString(s.C(tskv.FieldTs)).Comma().
						WriteString(strStartTs).
						WriteByte(')').
						String(),
					"ts_bucket",
				),
				//sql.As(
				//	builder.Reset().WriteString(strTimeBucket).String(),
				//	"interval",
				//),
				sql.As(sql.Max(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldLongV).Comma().
						WriteString("-9223372036854775807").
						WriteByte(')').
						String(),
				),
					"long_value",
				),
				sql.As(sql.Max(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldDblV).Comma().
						WriteString("-1.79769E+308").
						WriteByte(')').
						String(),
				),
					"double_value",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldLongV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"long_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldDblV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"double_value_Count",
				),
				sql.As(sql.Max(tskv.FieldStrV), "string_value"),
				//sql.As(`NULL`, "strValue"),
				//sql.As("MAX", "aggType"),
			).
				Where(
					sql.And(
						sql.EQ(tskv.FieldEntityID, entityId),
						sql.EQ(tskv.FieldKey, entityKeyId),
						sql.GT(tskv.FieldTs, startTs),
						sql.LTE(tskv.FieldTs, endTs),
					),
				).
				GroupBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket").
				OrderBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket")
		}).
		Scan(r.ctx, &results)
	return results, err
}

func (r *ThingsBoardRepository) FindMin(entityId uuid.UUID, entityKeyId int, timeBucket int64, startTs, endTs int64) ([]AggregationResult, error) {
	strTimeBucket := strconv.FormatInt(timeBucket, 10)
	strStartTs := strconv.FormatInt(startTs, 10)

	builder := &sql.Builder{}

	var results []AggregationResult

	err := r.client.Debug().TsKv.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(
					builder.Reset().WriteString("time_bucket(").
						WriteString(strTimeBucket).Comma().
						WriteString(s.C(tskv.FieldTs)).Comma().
						WriteString(strStartTs).
						WriteByte(')').
						String(),
					"ts_bucket",
				),
				//sql.As(
				//	builder.Reset().WriteString(strTimeBucket).String(),
				//	"interval",
				//),
				sql.As(sql.Min(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldLongV).Comma().
						WriteString("9223372036854775807").
						WriteByte(')').
						String(),
				),
					"long_value",
				),
				sql.As(sql.Min(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldDblV).Comma().
						WriteString("1.79769E+308").
						WriteByte(')').
						String(),
				),
					"double_value",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldLongV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"long_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldDblV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"double_value_Count",
				),
				sql.As(sql.Min(tskv.FieldStrV), "string_value"),
				//sql.As(`NULL`, "strValue"),
				//sql.As("MIN", "aggType"),
			).
				Where(
					sql.And(
						sql.EQ(tskv.FieldEntityID, entityId),
						sql.EQ(tskv.FieldKey, entityKeyId),
						sql.GT(tskv.FieldTs, startTs),
						sql.LTE(tskv.FieldTs, endTs),
					),
				).
				GroupBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket").
				OrderBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket")
		}).
		Scan(r.ctx, &results)
	return results, err
}

func (r *ThingsBoardRepository) FindSum(entityId uuid.UUID, entityKeyId int, timeBucket int64, startTs, endTs int64) ([]AggregationResult, error) {
	strTimeBucket := strconv.FormatInt(timeBucket, 10)
	strStartTs := strconv.FormatInt(startTs, 10)

	builder := &sql.Builder{}

	var results []AggregationResult

	err := r.client.Debug().TsKv.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(
					builder.Reset().WriteString("time_bucket(").
						WriteString(strTimeBucket).Comma().
						WriteString(s.C(tskv.FieldTs)).Comma().
						WriteString(strStartTs).
						WriteByte(')').
						String(),
					"ts_bucket",
				),
				//sql.As(
				//	builder.Reset().WriteString(strTimeBucket).String(),
				//	"interval",
				//),
				sql.As(sql.Sum(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldLongV).Comma().
						WriteString("0").
						WriteByte(')').
						String(),
				),
					"long_value",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("COALESCE(").
						Ident(tskv.FieldDblV).Comma().
						WriteString("0.0").
						WriteByte(')').
						String(),
				),
					"double_value",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldLongV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"long_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldDblV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"double_value_Count",
				),
				//sql.As(`NULL`, "strValue"),
				//sql.As("SUM", "aggType"),
			).
				Where(
					sql.And(
						sql.EQ(tskv.FieldEntityID, entityId),
						sql.EQ(tskv.FieldKey, entityKeyId),
						sql.GT(tskv.FieldTs, startTs),
						sql.LTE(tskv.FieldTs, endTs),
					),
				).
				GroupBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket").
				OrderBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket")
		}).
		Scan(r.ctx, &results)
	return results, err
}

func (r *ThingsBoardRepository) FindCount(entityId uuid.UUID, entityKeyId int, timeBucket int64, startTs, endTs int64) ([]AggregationResult, error) {
	strTimeBucket := strconv.FormatInt(timeBucket, 10)
	strStartTs := strconv.FormatInt(startTs, 10)

	builder := &sql.Builder{}

	var results []AggregationResult

	err := r.client.Debug().TsKv.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(
					builder.Reset().WriteString("time_bucket(").
						WriteString(strTimeBucket).Comma().
						WriteString(s.C(tskv.FieldTs)).Comma().
						WriteString(strStartTs).
						WriteByte(')').
						String(),
					"ts_bucket",
				),
				//sql.As(
				//	builder.Reset().WriteString(strTimeBucket).String(),
				//	"interval",
				//),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldBoolV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"bool_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldStrV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"string_value_Count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldLongV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"long_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldDblV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"double_value_count",
				),
				sql.As(sql.Sum(
					builder.Reset().WriteString("CASE WHEN").Pad().
						Ident(tskv.FieldJSONV).Pad().
						WriteString("IS NULL THEN 0 ELSE 1 END").
						String(),
				),
					"json_value_count",
				),
			).
				Where(
					sql.And(
						sql.EQ(tskv.FieldEntityID, entityId),
						sql.EQ(tskv.FieldKey, entityKeyId),
						sql.GT(tskv.FieldTs, startTs),
						sql.LTE(tskv.FieldTs, endTs),
					),
				).
				GroupBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket").
				OrderBy(tskv.FieldEntityID, tskv.FieldKey, "ts_bucket")
		}).
		Scan(r.ctx, &results)
	return results, err
}
