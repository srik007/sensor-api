package entity

import (
	"database/sql/driver"
	"encoding/json"

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

type Specie struct {
	Name  string
	Count int
}

type Species []Specie

func (s *Species) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &s)
}

func (s Species) Value() (driver.Value, error) {
	val, err := json.Marshal(s)
	return string(val), err
}

type SensorData struct {
	Transparency uint
	Temparature  valueobjects.Temparature `gorm:"embedded;"`
	Species      Species                  `gorm:"type:text;"`
	SensorId     uint                     `gorm:"index;primaryKey;"`
	Sensor       Sensor                   `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

func (s *Sensor) GetFakeData() SensorData {
	return GenerateFakeSensorData(*s)
}
