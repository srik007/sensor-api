package persistance

import (
	"fmt"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	Sensor      repository.SensorRepository
	SensorGroup repository.SensorGroupRepository
	db          *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Sensor:      NewSensorRepository(db),
		SensorGroup: NewSensorGroupRepository(db),
		db:          db,
	}, nil
}

func (s *Repositories) Close() error {
	sqlDb, err := s.db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDb.Close()
}

func (s *Repositories) Automigrate() error {
	if err := s.db.AutoMigrate(&entity.Sensor{}, &entity.SensorGroup{}, &entity.SensorData{}); err != nil {
		panic(err)
	}
	return nil
}
