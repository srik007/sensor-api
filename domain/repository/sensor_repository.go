package repository

import "github.com/srik007/sensor-api/domain/entity"

type SensorRepository interface {
	SaveAll(sensors []entity.Sensor) ([]entity.Sensor, map[string]string)
}
