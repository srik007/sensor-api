package scheduler

import (
	"time"

	"github.com/srik007/sensor-api/domain/entity"
)

func (j *SchedulerJob) ProduceSensorData(sensor entity.Sensor) {
	for {
		time.Sleep(time.Duration(sensor.DataOutputRate.Value) * time.Second)
		sensorData := sensor.GetFakeData()
		select {
		case j.Channles[int(sensor.ID)] <- sensorData:
		}
	}
}
