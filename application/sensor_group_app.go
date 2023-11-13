package application

import (
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
)

type SensorGroupApp struct {
	repository repository.SensorGroupRepository
}

func (sg *SensorGroupApp) SaveAll(sensorGroups []entity.SensorGroup) ([]entity.SensorGroup, map[string]string) {
	return sg.repository.SaveAll(sensorGroups)
}

var _ SensorGroupAppInterface = &SensorGroupApp{}

type SensorGroupAppInterface interface {
	SaveAll(sensorGroups []entity.SensorGroup) ([]entity.SensorGroup, map[string]string)
}
