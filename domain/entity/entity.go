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

type Specie struct {
	Name  string
	Count int
}

type SensorData struct {
	Transparency int
	Temparature  valueobjects.Temparature `gorm:"embedded;"`
	Specie       []Specie                 `gorm:"embedded;"`
	Sensor       Sensor                   `gorm:"foreignKey:SensorId"`
}

func (s *Sensor) AfterCreate(tx *gorm.DB) (err error) {
	sensorGroup := &SensorGroup{
		Name:        s.CodeName.Name,
		SensorCount: int(s.CodeName.GroupId),
	}
	tx.Save(&sensorGroup)
	return nil
}
