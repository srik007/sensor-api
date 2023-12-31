package persistance

import (
	"fmt"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	"gorm.io/gorm"
)

type SensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) *SensorRepository {

	return &SensorRepository{db: db}
}

var _ repository.SensorRepository = &SensorRepository{}

func (r *SensorRepository) SaveAll(sensors []entity.Sensor) ([]entity.Sensor, map[string]string) {
	var result []entity.Sensor
	for _, sensor := range sensors {
		var existingSensorGroup entity.SensorGroup
		if err := r.db.Where("name = ?", sensor.CodeName.Name).First(&existingSensorGroup).Error; err != nil {
			panic("Configure sensor groups: " + err.Error())
		} else {
			sensor.CodeName.GroupId = uint64(existingSensorGroup.SensorCount) + 1
		}
		tx := r.db.Create(&sensor)
		if tx.Error != nil {
			panic(tx.Error)
		}
		result = append(result, sensor)
	}
	return result, nil
}

func (r *SensorRepository) GetAll() []entity.Sensor {
	var sensors []entity.Sensor
	if err := r.db.Find(&sensors).Error; err != nil {
		fmt.Printf("Error retrieving products: " + err.Error())
	}
	return sensors
}

func (r *SensorRepository) SaveData(sensorData entity.SensorData) {
	if err := r.db.Save(&sensorData).Error; err != nil {
		fmt.Errorf("Failed to update sensor data")
	}
}
