package application

import (
	"github.com/srik007/sensor-api/domain/aggregators"
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/generators"
	"github.com/srik007/sensor-api/domain/repository"
	scheduler "github.com/srik007/sensor-api/domain/schedulers"
	"gorm.io/gorm"
)

type SensorApp struct {
	Generator  generators.Generator
	Scheduler  scheduler.SchedulerJob
	Aggregator aggregators.Aggregator
}

func NewSensorApp(s repository.SensorRepository, sg repository.SensorGroupRepository, dataStore *gorm.DB) *SensorApp {
	return &SensorApp{
		Generator:  *generators.NewGenerator(s, sg),
		Scheduler:  *scheduler.NewJob(s),
		Aggregator: *aggregators.NewAggregator(dataStore),
	}
}

type SensorAppInterface interface {
	GenerateMetadata()
	ScheduleJob()
	CollectSpeciesUnderGroup(groupName string) entity.Species
	CollectTopNSpeciesUnderGroup(groupName string, topN int) entity.Species
}

var _ SensorAppInterface = &SensorApp{}

func (s *SensorApp) GenerateMetadata() {
	s.Generator.GenerateSensorMetaData()
	s.Generator.GenerateSensorGroupMetaData()
}

func (s *SensorApp) ScheduleJob() {
	s.Scheduler.Run()
	s.Aggregator.Run()
}

func (s *SensorApp) CollectSpeciesUnderGroup(groupName string) entity.Species {
	return s.Aggregator.CollectSpeciesUnderGroup(groupName)
}

func (s *SensorApp) CollectTopNSpeciesUnderGroup(groupName string, topN int) entity.Species {
	return s.Aggregator.CollectTopNSpeciesUnderGroup(groupName, topN)
}
