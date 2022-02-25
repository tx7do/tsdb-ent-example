package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

// AttributeKv 属性键值
type AttributeKv struct {
	EntityType          string    `json:"entity_type" bson:"entity_type"`
	EntityId            uuid.UUID `json:"entity_id" bson:"entity_id"`
	AttributeType       string    `json:"attribute_type" bson:"attribute_type"`
	AttributeKey        string    `json:"attribute_key" bson:"attribute_key"`
	LastUpdateTimestamp int64     `json:"last_update_ts" bson:"last_update_ts"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

// TsKv 遥测数据 历史数据 (时序数据)
type TsKv struct {
	EntityId  uuid.UUID `json:"entity_id" bson:"entity_id"`
	Key       int       `json:"key" bson:"key"`
	Timestamp int64     `json:"ts" bson:"ts"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

func NewTsKv(entityId uuid.UUID, keyId int, timestamp int64,
	boolV *bool, stringV *string, longV *int64, doubleV *float64, jsonV *string) *TsKv {
	c := &TsKv{
		EntityId:  entityId,
		Key:       keyId,
		Timestamp: timestamp,

		BoolV:   boolV,
		StringV: stringV,
		LongV:   longV,
		DoubleV: doubleV,
		JsonV:   jsonV,
	}
	return c
}

// TsKvDictionary 键值字典
type TsKvDictionary struct {
	KeyId   int    `json:"key_id" bson:"_id"`      // 编号
	Key     string `json:"key" bson:"key"`         // 键名,如: humidity
	Display string `json:"display" bson:"display"` // 显示名,如: 湿度
	Unit    string `json:"unit" bson:"unit"`       // 单位,如: RH
}

// TsKvLatest 遥测数据 最新数据
type TsKvLatest struct {
	EntityId  uuid.UUID `json:"entity_id" bson:"entity_id"`
	Key       int       `json:"key" bson:"key"`
	Timestamp int64     `json:"ts" bson:"ts"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

func NewTsKvLatest(entityId uuid.UUID, keyId int, timestamp int64,
	boolV *bool, stringV *string, longV *int64, doubleV *float64, jsonV *string) *TsKvLatest {
	c := &TsKvLatest{
		EntityId:  entityId,
		Key:       keyId,
		Timestamp: timestamp,

		BoolV:   boolV,
		StringV: stringV,
		LongV:   longV,
		DoubleV: doubleV,
		JsonV:   jsonV,
	}
	return c
}

type KvMap struct {
	Key string `json:"key" bson:"key"`

	BoolV   *bool    `json:"bool_v,omitempty" bson:"bool_v,omitempty"`
	StringV *string  `json:"str_v,omitempty" bson:"str_v,omitempty"`
	LongV   *int64   `json:"long_v,omitempty" bson:"long_v,omitempty"`
	DoubleV *float64 `json:"dbl_v,omitempty" bson:"dbl_v,omitempty"`
	JsonV   *string  `json:"json_v,omitempty" bson:"json_v,omitempty"`
}

func MarshalKv(value interface{}) []KvMap {
	th := reflect.TypeOf(value)
	vh := reflect.ValueOf(value)

	var result KvMap
	var results = make([]KvMap, 0)
	for i := 0; i < th.NumField(); i++ {
		fieldType := th.Field(i)
		fieldValue := vh.FieldByName(fieldType.Name)

		tagJson := fieldType.Tag.Get("json")
		if idx := strings.Index(tagJson, ","); idx != -1 {
			tagJson = tagJson[:idx]
		}

		result.Key = tagJson
		result.BoolV = nil
		result.StringV = nil
		result.LongV = nil
		result.DoubleV = nil
		result.JsonV = nil

		switch fieldType.Type.Kind() {
		case reflect.Bool:
			v := fieldValue.Interface().(bool)
			results = append(results, KvMap{Key: tagJson, BoolV: &v})
			
		case reflect.String:
			v := fieldValue.Interface().(string)
			results = append(results, KvMap{Key: tagJson, StringV: &v})

		case reflect.Int:
			v := int64(fieldValue.Interface().(int))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Int8:
			v := int64(fieldValue.Interface().(int8))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Int16:
			v := int64(fieldValue.Interface().(int16))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Int32:
			v := int64(fieldValue.Interface().(int32))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Int64:
			v := fieldValue.Interface().(int64)
			results = append(results, KvMap{Key: tagJson, LongV: &v})

		case reflect.Uint:
			v := int64(fieldValue.Interface().(uint))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Uint8:
			v := int64(fieldValue.Interface().(uint8))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Uint16:
			v := int64(fieldValue.Interface().(uint16))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Uint32:
			v := int64(fieldValue.Interface().(uint32))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Uintptr:
			v := int64(fieldValue.Interface().(uintptr))
			results = append(results, KvMap{Key: tagJson, LongV: &v})
		case reflect.Uint64:
			v := int64(fieldValue.Interface().(uint64))
			results = append(results, KvMap{Key: tagJson, LongV: &v})

		case reflect.Float32:
			v := float64(fieldValue.Interface().(float32))
			results = append(results, KvMap{Key: tagJson, DoubleV: &v})
		case reflect.Float64:
			v := fieldValue.Interface().(float64)
			results = append(results, KvMap{Key: tagJson, DoubleV: &v})

		case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
			jsonByte, _ := json.Marshal(fieldValue.Interface())
			v := string(jsonByte)
			results = append(results, KvMap{Key: tagJson, JsonV: &v})
		}
	}

	return results
}
