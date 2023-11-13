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
	fmt.Println(sensorGroups)
	for _, sensorGroup := range sensorGroups {
		result := r.db.Create(&sensorGroup)
		if result.Error != nil {
			panic(result.Error)
		}
		fmt.Println("SensorGroup created successfully:", sensorGroup)
	}
	return sensorGroups, nil
}

// func (r *SensorGroupRepository) UpdateSensorCount(name string, count int) {
// 	if err := r.db.Model(&entity.SensorGroup{}).Where("name = ?", name).Update("sensor_count", count); err != nil {
// 		panic("Error updating sensor count")
// 	}
// }
