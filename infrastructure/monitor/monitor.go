package monitor

import (
	"fmt"
	"sync"
	"time"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
)

func Register(sensors []entity.Sensor) {
	channels := make(map[int]chan entity.SensorData)
	var mutex sync.Mutex
	for _, sensor := range sensors {
		channel := make(chan entity.SensorData)
		mutex.Lock()
		channels[int(sensor.ID)] = channel
		mutex.Unlock()
	}
	for _, sensor := range sensors {
		//Consumer
		go func(channel chan entity.SensorData) {
			for {
				event := <-channel
				fmt.Println("Hey coming here", event)
			}
		}(channels[int(sensor.ID)])

		//Producer
		go func(sensor entity.Sensor) {
			for {
				time.Sleep(3 * time.Second)
				sensorData := &entity.SensorData{
					Temparature: valueObjects.Temparature{Value: 38, Scale: "Celsius"},
					Specie: []entity.Specie{
						{Name: "Atlantic Cod",
							Count: 12}, {Name: "Sailfish", Count: 4},
					},
					Sensor:       sensor,
					Transparency: 10,
				}
				select {
				case channels[int(sensor.ID)] <- *sensorData:
				}
			}
		}(sensor)
	}
}
