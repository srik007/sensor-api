package scheduler

import (
	"fmt"
	"sync"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
)

type SchedulerJob struct {
	SensorRepository repository.SensorRepository
	Channles         map[int]chan entity.SensorData
}

func NewJob(s repository.SensorRepository) *SchedulerJob {
	return &SchedulerJob{SensorRepository: s, Channles: make(map[int]chan entity.SensorData)}
}

func (j *SchedulerJob) Run() {
	sensors := j.SensorRepository.GetAll()
	if len(sensors) == 0 {
		fmt.Printf("No sensors are available to schedule.")
	}
	var mutex sync.Mutex
	for _, sensor := range sensors {
		channel := make(chan entity.SensorData)
		mutex.Lock()
		j.Channles[int(sensor.ID)] = channel
		mutex.Unlock()
	}
	for _, sensor := range sensors {
		go j.RecieveSensorData(int(sensor.ID))
		go j.ProduceSensorData(sensor)
	}
}
