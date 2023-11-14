package monitor

import (
	"time"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
)

func (j *SensorDataMonitorJob) ProduceSensorData(sensor entity.Sensor) {
	for {
		time.Sleep(time.Duration(sensor.DataOutputRate.Value) * time.Second)
		sensorData := []entity.SensorData{{
			Temparature: valueObjects.Temparature{Value: 38, Scale: "Celsius"},
			Species: []entity.Specie{
				{Name: "Atlantic Cod",
					Count: 12}, {Name: "Sailfish", Count: 4},
			},
			SensorId:     uint(sensor.ID),
			Transparency: 10,
		}}
		select {
		case j.Channles[int(sensor.ID)] <- sensorData:
		}
	}
}
