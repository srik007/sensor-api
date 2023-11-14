package application

import (
	"github.com/srik007/sensor-api/domain/generators"
	"github.com/srik007/sensor-api/domain/repository"
	scheduler "github.com/srik007/sensor-api/domain/schedulers"
)

type SensorApp struct {
	SensorRepository      repository.SensorRepository
	SensorGroupRepository repository.SensorGroupRepository
}

type SensorAppInterface interface {
	GenerateMetadata()
	ScheduleJob()
}

var _ SensorAppInterface = &SensorApp{}

func (s *SensorApp) GenerateMetadata() {
	generator := generators.NewGenerator(s.SensorRepository, s.SensorGroupRepository)
	generator.GenerateSensorMetaData()
	generator.GenerateSensorGroupMetaData()
}

func (s *SensorApp) ScheduleJob() {
	scheduler := scheduler.NewJob(s.SensorRepository)
	scheduler.Run()
}
