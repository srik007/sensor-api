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
}

var _ SensorAppInterface = &SensorApp{}

func (s *SensorApp) SaveAll(sensors []entity.Sensor) ([]entity.Sensor, map[string]string) {
	return s.repository.SaveAll(sensors)
}
