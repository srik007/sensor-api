package persistance

import (
	"fmt"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	"gorm.io/gorm"
)

type SensorGroupRepository struct {
	db *gorm.DB
}

func NewSensorGroupRepository(db *gorm.DB) *SensorGroupRepository {
	return &SensorGroupRepository{db: db}
}

var _ repository.SensorGroupRepository = &SensorGroupRepository{}

func (r *SensorGroupRepository) SaveAll(sensorGroups []entity.SensorGroup) ([]entity.SensorGroup, map[string]string) {
	for _, sensorGroup := range sensorGroups {
		r.db.Create(&sensorGroup)
		fmt.Println("SensorGroup created successfully:", sensorGroup)
	}
	return sensorGroups, nil
}
