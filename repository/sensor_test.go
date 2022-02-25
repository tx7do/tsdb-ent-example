package repository

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"tsdb-ent-example/model"
	"tsdb-ent-example/util"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestInsertSensorData(t *testing.T) {
	client := NewSensorRepository()
	assert.Nil(t, client.InsertSensor(&model.Sensor{Type: "a", Location: "floor"}))
	assert.Nil(t, client.InsertSensor(&model.Sensor{Type: "a", Location: "ceiling"}))
	assert.Nil(t, client.InsertSensor(&model.Sensor{Type: "b", Location: "floor"}))
	assert.Nil(t, client.InsertSensor(&model.Sensor{Type: "b", Location: "ceiling"}))
}

func TestInsertSensorTimelineData(t *testing.T) {
	client := NewSensorRepository()

	var datas []*model.SensorData

	nowTime := time.Now()
	startTime := nowTime
	startTime = startTime.Add(-time.Hour * 24)
	for {
		startTime = startTime.Add(time.Minute * 5)

		if startTime.Unix() >= nowTime.Unix() {
			break
		}

		datas = append(datas, &model.SensorData{
			Time:        startTime,
			SensorId:    util.RandomInt(1, 4),
			CPU:         util.RandomFloat(1, 100),
			Temperature: util.RandomFloat(0, 90)})
	}
	assert.Nil(t, client.BatchInsertSensorData(datas))
}

func TestAggregation1(t *testing.T) {
	client := NewSensorRepository()
	v, err := client.Avg()
	assert.Nil(t, err)
	fmt.Println(v)
}

func TestAggregation2(t *testing.T) {
	client := NewSensorRepository()
	v, err := client.AvgAndLatest()
	assert.Nil(t, err)
	fmt.Println(v)
}
