package repository

import "github.com/srik007/sensor-api/domain/entity"

type SensorGroupRepository interface {
	SaveAll(sensorGroups []entity.SensorGroup) ([]entity.SensorGroup, map[string]string)
	//UpdateSensorCount(name string, value int)
}
