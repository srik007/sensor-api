package monitor

import (
	"sync"

	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/entity"
)

type SensorDataMonitorJob struct {
	SensorApp application.SensorAppInterface
	Channles  map[int]chan []entity.SensorData
}

func NewJob(sapp application.SensorAppInterface) *SensorDataMonitorJob {
	return &SensorDataMonitorJob{SensorApp: sapp, Channles: make(map[int]chan []entity.SensorData)}
}

func (j *SensorDataMonitorJob) Register(sensors []entity.Sensor) {
	var mutex sync.Mutex
	for _, sensor := range sensors {
		channel := make(chan []entity.SensorData)
		mutex.Lock()
		j.Channles[int(sensor.ID)] = channel
		mutex.Unlock()
	}
	for _, sensor := range sensors {
		go j.RecieveSensorData(int(sensor.ID))
		go j.ProduceSensorData(sensor)
	}
}
