package persistance

import (
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
	for _, sensor := range sensors {
		var existingSensorGroup entity.SensorGroup
		if err := r.db.Where("name = ?", sensor.CodeName.Name).First(&existingSensorGroup).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				sensor.CodeName.GroupId = 1
			} else {
				panic("Error retrieving sensor: " + err.Error())
			}
		} else {
			sensor.CodeName.GroupId = uint64(existingSensorGroup.SensorCount) + 1
		}
		result := r.db.Create(&sensor)
		if result.Error != nil {
			panic(result.Error)
		}
	}
	return sensors, nil
}
