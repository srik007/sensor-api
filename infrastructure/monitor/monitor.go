package monitor

import (
	"fmt"
	"sync"
	"time"

	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
)

type SensorDataMonitorJob struct {
	SensorApp application.SensorAppInterface
}

func (j *SensorDataMonitorJob) Register(sensors []entity.Sensor) {
	channels := make(map[int]chan []entity.SensorData)
	var mutex sync.Mutex
	for _, sensor := range sensors {
		channel := make(chan []entity.SensorData)
		mutex.Lock()
		channels[int(sensor.ID)] = channel
		mutex.Unlock()
	}
	for _, sensor := range sensors {
		go HandleSensorData(channels[int(sensor.ID)], j.SensorApp)
		go ProduceSensorData(channels[int(sensor.ID)], sensor)
	}
}

func HandleSensorData(channel chan []entity.SensorData, sapp application.SensorAppInterface) {
	for {
		event := <-channel
		sapp.SaveData(event)
	}
}

func ProduceSensorData(channel chan []entity.SensorData, sensor entity.Sensor) {
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
		fmt.Printf("%v", sensorData[0])
		select {
		case channel <- sensorData:
		}
	}
}
