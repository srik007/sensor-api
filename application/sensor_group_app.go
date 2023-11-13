package application

import (
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
)

type SensorGroupAppInterface interface {
	SaveAll(sensorGroups []entity.SensorGroup) ([]entity.SensorGroup, map[string]string)
	//UpdateSensorCount(name string, value int)
}

type SensorGroupApp struct {
	repository repository.SensorGroupRepository
}

func (s *SensorGroupApp) SaveAll(sensorGroups []entity.SensorGroup) ([]entity.SensorGroup, map[string]string) {
	return s.repository.SaveAll(sensorGroups)
}
