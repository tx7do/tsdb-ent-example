package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"tsdb-ent-example/ent"
	"tsdb-ent-example/ent/sensordata"
	"tsdb-ent-example/model"
)

type SensorAvgData struct {
	Timestamp      time.Time `json:"period"`
	AvgTemperature float64   `json:"avg_temp"`
	AvgCpu         float64   `json:"avg_cpu"`
}

type SensorAvgAndLatestData struct {
	Timestamp       time.Time `json:"period"`
	AvgTemperature  float64   `json:"avg_temp"`
	AvgCpu          float64   `json:"avg_cpu"`
	LastTemperature float64   `json:"last_temp"`
}

type SensorRepository struct {
	client *ent.Client
	ctx    context.Context
}

func NewSensorRepository() *SensorRepository {
	r := &SensorRepository{}
	_ = r.Connect()
	return r
}

func (r *SensorRepository) Connect() error {
	cli, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=test sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// Run the auto migration tool.
	//if err := cli.Schema.Create(context.Background()); err != nil {
	//	fmt.Printf("failed creating schema resources: %v \n", err)
	//}

	r.ctx = context.Background()
	r.client = cli

	return err
}

func (r *SensorRepository) InsertSensor(s *model.Sensor) error {
	return r.client.Sensor.Create().
		SetType(s.Type).
		SetLocation(s.Location).
		Exec(r.ctx)
}

func (r *SensorRepository) InsertSensorData(s *model.SensorData) error {
	return r.client.SensorData.Create().
		SetTime(s.Time).
		SetSensorID(s.SensorId).
		SetCPU(s.CPU).
		SetTemperature(s.Temperature).
		Exec(r.ctx)
}

func (r *SensorRepository) BatchInsertSensorData(sd []*model.SensorData) error {
	bulks := make([]*ent.SensorDataCreate, 0)
	for i := 0; i < len(sd); i++ {
		s := sd[i]
		bulk := r.client.SensorData.Create().
			SetTime(s.Time).
			SetSensorID(s.SensorId).
			SetCPU(s.CPU).
			SetTemperature(s.Temperature)
		bulks = append(bulks, bulk)
	}
	return r.client.SensorData.CreateBulk(bulks...).Exec(r.ctx)
}

func (r *SensorRepository) Avg() ([]SensorAvgData, error) {
	var v []SensorAvgData
	err := r.client.Debug().SensorData.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As("time_bucket('30 minutes', time)", "period"),
				sql.As(sql.Avg(sensordata.FieldTemperature), "avg_temp"),
				sql.As(sql.Avg(sensordata.FieldCPU), "avg_cpu"),
			).
				GroupBy("period")
		}).
		Scan(r.ctx, &v)
	return v, err
}

func (r *SensorRepository) AvgAndLatest() ([]SensorAvgAndLatestData, error) {
	var v []SensorAvgAndLatestData
	err := r.client.Debug().SensorData.Query().
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As("time_bucket('30 minutes', time)", "period"),
				sql.As(sql.Avg(sensordata.FieldTemperature), "avg_temp"),
				sql.As(sql.Avg(sensordata.FieldCPU), "avg_cpu"),
				sql.As("last(temperature, time)", "last_temp"),
			).
				GroupBy("period")
		}).
		Scan(r.ctx, &v)
	return v, err
}
