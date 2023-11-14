package application

import (
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
)

type SensorApp struct {
	repository repository.SensorRepository
}

type SensorAppInterface interface {
	SaveAll(sensors []entity.Sensor) ([]entity.Sensor, map[string]string)
	GetAll() []entity.Sensor
	SaveData(sensorData entity.SensorData)
}

var _ SensorAppInterface = &SensorApp{}

func (s *SensorApp) SaveAll(sensors []entity.Sensor) ([]entity.Sensor, map[string]string) {
	return s.repository.SaveAll(sensors)
}

func (s *SensorApp) GetAll() []entity.Sensor {
	return s.repository.GetAll()
}

func (s *SensorApp) SaveData(sensorData entity.SensorData) {
	s.repository.SaveData(sensorData)
}
