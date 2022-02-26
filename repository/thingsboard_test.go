package repository

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
	"tsdb-ent-example/model"
	"tsdb-ent-example/util"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Gps GPS设备数据
type Gps struct {
	Latitude        float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`               // 纬度
	Longitude       float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`             // 经度
	BatteryLevel    int     `json:"batteryLevel,omitempty" bson:"batteryLevel,omitempty"`       // 电池电量
	BatteryCharging bool    `json:"batteryCharging,omitempty" bson:"batteryCharging,omitempty"` // 电池充电状态
}

// Hygrothermograph 温湿度计数据
type Hygrothermograph struct {
	Humidity    float64 `json:"humidity,omitempty" bson:"humidity,omitempty"`       // 湿度
	Temperature float64 `json:"temperature,omitempty" bson:"temperature,omitempty"` // 温度
}

func TestSaveKeyId(t *testing.T) {
	client := NewThingsBoardRepository()

	assert.Nil(t, client.SaveKeyId("humidity"))
	assert.Nil(t, client.SaveKeyId("temperature"))

	assert.Nil(t, client.SaveKeyId("latitude"))
	assert.Nil(t, client.SaveKeyId("longitude"))
	assert.Nil(t, client.SaveKeyId("batteryLevel"))
	assert.Nil(t, client.SaveKeyId("batteryCharging"))
}

func TestGetOrSaveKeyId(t *testing.T) {
	client := NewThingsBoardRepository()

	assert.NotEqual(t, client.GetOrSaveKeyId("humidity"), -1)
	assert.NotEqual(t, client.GetOrSaveKeyId("temperature"), -1)

	assert.NotEqual(t, client.GetOrSaveKeyId("latitude"), -1)
	assert.NotEqual(t, client.GetOrSaveKeyId("longitude"), -1)
	assert.NotEqual(t, client.GetOrSaveKeyId("batteryLevel"), -1)
	assert.NotEqual(t, client.GetOrSaveKeyId("batteryCharging"), -1)
}

func parseTag(tag string) string {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx]
	}
	return tag
}

func TestReflect(t *testing.T) {
	type Sub struct {
		A int `json:"a,omitempty" bson:"a,omitempty"`
		B int `json:"b,omitempty" bson:"b,omitempty"`
	}
	type ReflectStruct struct {
		Humidity    float64 `json:"humidity,omitempty" bson:"humidity,omitempty"`       // 湿度
		Temperature float64 `json:"temperature,omitempty" bson:"temperature,omitempty"` // 温度

		Sub Sub `json:"sub,omitempty" bson:"sub,omitempty"`
	}

	hygrothermograph := ReflectStruct{
		Humidity:    util.RandomFloat(0, 90),
		Temperature: util.RandomFloat(-40, 60),
		Sub:         Sub{A: 1, B: 2},
	}

	{
		results := model.MarshalKv(hygrothermograph)
		assert.Equal(t, len(results), 3)
	}

	{
		th := reflect.TypeOf(hygrothermograph)
		vh := reflect.ValueOf(hygrothermograph)

		for i := 0; i < th.NumField(); i++ {
			fieldType := th.Field(i)
			fieldValue := vh.FieldByName(fieldType.Name)

			tagJson := fieldType.Tag.Get("json")
			if idx := strings.Index(tagJson, ","); idx != -1 {
				tagJson = tagJson[:idx]
			}

			fmt.Printf("name: %v  json tag: '%v'\n", fieldType.Name, tagJson)

			fmt.Println(fieldType.Name, fieldType.Type.Kind(), fieldValue.Interface())

			switch fieldType.Type.Kind() {
			case reflect.Bool:
			case reflect.String:
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			case reflect.Float32, reflect.Float64:
				v := fieldValue.Interface()
				fmt.Println(v)
			case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
				jsonByte, _ := json.Marshal(fieldValue.Interface())
				fmt.Println(string(jsonByte))
			}

		}
	}
}

func TestSaveOrUpdateTsKv(t *testing.T) {
	client := NewThingsBoardRepository()

	{
		entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
		assert.Nil(t, err)

		timestamp := time.Now().UnixNano()

		hygrothermograph := Hygrothermograph{
			Humidity:    util.RandomFloat(0, 90),
			Temperature: util.RandomFloat(-40, 60),
		}

		results := model.MarshalKv(hygrothermograph)

		for k, v := range results {
			tsKv := model.NewTsKv(entityId, client.GetOrSaveKeyId(k), timestamp,
				v.BoolV, v.StringV, v.LongV, v.DoubleV, v.JsonV)
			err := client.SaveOrUpdateTsKv(tsKv)
			assert.Nil(t, err)
			err = client.SaveOrUpdateTsKvLatest(tsKv)
			assert.Nil(t, err)
		}
	}

	{
		entityId, err := uuid.Parse("ad34ff10-7514-11ec-9a90-af0223be0666")
		assert.Nil(t, err)

		timestamp := time.Now().UnixNano()

		//全球经纬度的取值范围为：纬度-90至90，经度-180至180
		//中国的经纬度范围大约为：纬度3.86至53.55，经度73.66至135.05
		//北京行政中心的纬度为39.92，经度为116.46

		gps := Gps{
			Latitude:        util.RandomFloat(3.86, 53.55),
			Longitude:       util.RandomFloat(73.66, 135.05),
			BatteryLevel:    util.RandomInt(0, 100),
			BatteryCharging: true,
		}

		results := model.MarshalKv(gps)

		for k, v := range results {
			tsKv := model.NewTsKv(entityId, client.GetOrSaveKeyId(k), timestamp,
				v.BoolV, v.StringV, v.LongV, v.DoubleV, v.JsonV)
			err := client.SaveOrUpdateTsKv(tsKv)
			assert.Nil(t, err)
			err = client.SaveOrUpdateTsKvLatest(tsKv)
			assert.Nil(t, err)
		}
	}
}

func TestSaveOrUpdateTsKvLatest(t *testing.T) {
	//client := NewThingsBoardRepository()
}

func TestFindTsKvLatest(t *testing.T) {
	client := NewThingsBoardRepository()

	entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)

	results, err := client.FindTsKvLatest(entityId)
	assert.Nil(t, err)

	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}
}

func TestFindAllKeysByEntityIds(t *testing.T) {
	client := NewThingsBoardRepository()

	var entityIds []interface{}

	entityId1, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)
	entityIds = append(entityIds, entityId1)

	entityId2, err := uuid.Parse("ad34ff10-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)
	entityIds = append(entityIds, entityId2)

	results, err := client.FindAllKeysByEntityIds(entityIds)
	assert.Nil(t, err)
	fmt.Printf("%+v\n", results)
}

func TestFindAvg(t *testing.T) {
	client := NewThingsBoardRepository()

	entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)

	startTime, _ := time.Parse(`2006-01-02 15:04:05`, `2021-01-01 00:00:00`)
	endTime := time.Now()
	results, err := client.FindAvg(entityId, 1, int64(time.Second*30), startTime.UnixNano(), endTime.UnixNano())
	assert.Nil(t, err)
	fmt.Printf("%+v\n", results)
}

func TestFindMax(t *testing.T) {
	client := NewThingsBoardRepository()

	entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)

	startTime, _ := time.Parse(`2006-01-02 15:04:05`, `2021-01-01 00:00:00`)
	endTime := time.Now()
	results, err := client.FindMax(entityId, 1, int64(time.Second*30), startTime.UnixNano(), endTime.UnixNano())
	assert.Nil(t, err)
	fmt.Printf("%+v\n", results)
}

func TestFindMin(t *testing.T) {
	client := NewThingsBoardRepository()

	entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)

	startTime, _ := time.Parse(`2006-01-02 15:04:05`, `2021-01-01 00:00:00`)
	endTime := time.Now()
	results, err := client.FindMin(entityId, 1, int64(time.Second*30), startTime.UnixNano(), endTime.UnixNano())
	assert.Nil(t, err)
	fmt.Printf("%+v\n", results)
}

func TestFindSum(t *testing.T) {
	client := NewThingsBoardRepository()

	entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)

	startTime, _ := time.Parse(`2006-01-02 15:04:05`, `2021-01-01 00:00:00`)
	endTime := time.Now()
	results, err := client.FindSum(entityId, 1, int64(time.Second*30), startTime.UnixNano(), endTime.UnixNano())
	assert.Nil(t, err)
	fmt.Printf("%+v\n", results)
}

func TestFindCount(t *testing.T) {
	client := NewThingsBoardRepository()

	entityId, err := uuid.Parse("ad2bfe60-7514-11ec-9a90-af0223be0666")
	assert.Nil(t, err)

	startTime, _ := time.Parse(`2006-01-02 15:04:05`, `2021-01-01 00:00:00`)
	endTime := time.Now()
	results, err := client.FindCount(entityId, 1, int64(time.Second*30), startTime.UnixNano(), endTime.UnixNano())
	assert.Nil(t, err)
	fmt.Printf("%+v\n", results)
}
