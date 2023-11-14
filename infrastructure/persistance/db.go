package persistance

import (
	"fmt"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	SensorRepository      repository.SensorRepository
	SensorGroupRepository repository.SensorGroupRepository
	DataStore             *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		SensorRepository:      NewSensorRepository(db),
		SensorGroupRepository: NewSensorGroupRepository(db),
		DataStore:             db,
	}, nil
}

func (s *Repositories) Close() error {
	sqlDb, err := s.DataStore.DB()
	if err != nil {
		panic(err)
	}
	return sqlDb.Close()
}

func (s *Repositories) Automigrate() error {
	if err := s.DataStore.AutoMigrate(&entity.Sensor{}, &entity.SensorGroup{}, &entity.SensorData{}); err != nil {
		panic(err)
	}
	return nil
}
