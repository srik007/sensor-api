package entity

import (
	valueobjects "github.com/srik007/sensor-api/domain/valueObjects"
	"gorm.io/gorm"
)

type SensorGroup struct {
	Name        string `gorm:"unique;primaryKey;"`
	SensorCount int
}

type CodeName struct {
	Name    string `gorm:"index:idx_member"`
	GroupId uint64 `gorm:"index:idx_member;"`
}

type Sensor struct {
	ID             uint                       `gorm:"primaryKey;"`
	CodeName       CodeName                   `gorm:"embedded;"`
	Coordinate     valueobjects.Coordinate    `gorm:"embedded"`
	DataOutputRate valueobjects.DataOuputRate `gorm:"embedded"`
}

func (s *Sensor) AfterCreate(tx *gorm.DB) (err error) {
	sensorGroup := &SensorGroup{
		Name:        s.CodeName.Name,
		SensorCount: int(s.CodeName.GroupId),
	}
	tx.Save(&sensorGroup)
	return nil
}
