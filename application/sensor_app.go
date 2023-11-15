package application

import (
	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/aggregators"
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	scheduler "github.com/srik007/sensor-api/domain/schedulers"
	sensormetadata "github.com/srik007/sensor-api/domain/sensor_metadata"
	"github.com/srik007/sensor-api/domain/valueObjects"
	"github.com/srik007/sensor-api/infrastructure/cache"
	"gorm.io/gorm"
)

type SensorApp struct {
	MetadataCreator       sensormetadata.MetadataCreator
	Scheduler             scheduler.SchedulerJob
	AggregateQueryHandler aggregators.AggregatorQueryHandler
	AggregateScheduler    aggregators.AggregatorSchedulerJob
}

func NewSensorApp(s repository.SensorRepository, sg repository.SensorGroupRepository, dataStore *gorm.DB, cache *cache.RedisCache) *SensorApp {
	return &SensorApp{
		MetadataCreator:       *sensormetadata.NewMetadataCreator(s, sg),
		Scheduler:             *scheduler.NewScheduler(s),
		AggregateScheduler:    *aggregators.NewScheduler(cache),
		AggregateQueryHandler: *aggregators.NewAggregatorQueryHandler(dataStore),
	}
}

type SensorAppInterface interface {
	CreateMetadata()
	Schedule(c *gin.Context)
	CollectSpeciesUnderGroup(groupName string) entity.Species
	CollectTopNSpeciesUnderGroup(groupName string, topN int) entity.Species
	CollectAverageTransparencyUnderGroup(groupName string) float64
	CollectAverageTemparatureUnderGroup(groupName string) float64
	CalculateMinTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax float64) valueObjects.Temparature
	CalculateMaxTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax float64) valueObjects.Temparature
}

var _ SensorAppInterface = &SensorApp{}

func (s *SensorApp) CreateMetadata() {
	s.MetadataCreator.CreateSensorMetadata()
	s.MetadataCreator.CreateSensorGroupMetaData()
}

func (s *SensorApp) Schedule(c *gin.Context) {
	s.Scheduler.Run()
	s.AggregateScheduler.Run(c, s.AggregateQueryHandler)
}

func (s *SensorApp) CollectSpeciesUnderGroup(groupName string) entity.Species {
	return s.AggregateQueryHandler.CollectSpeciesUnderGroup(groupName)
}

func (s *SensorApp) CollectTopNSpeciesUnderGroup(groupName string, topN int) entity.Species {
	return s.AggregateQueryHandler.CollectTopNSpeciesUnderGroup(groupName, topN)
}

func (s *SensorApp) CollectAverageTransparencyUnderGroup(groupName string) float64 {
	return s.AggregateQueryHandler.CollectGroupAggregatorsByName(groupName).AverageTransparency
}

func (s *SensorApp) CollectAverageTemparatureUnderGroup(groupName string) float64 {
	return s.AggregateQueryHandler.CollectGroupAggregatorsByName(groupName).AverageTemperature
}

func (s *SensorApp) CalculateMinTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax float64) valueObjects.Temparature {
	return s.AggregateQueryHandler.CalculateMinTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax)
}

func (s *SensorApp) CalculateMaxTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax float64) valueObjects.Temparature {
	return s.AggregateQueryHandler.CalculateMaxTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax)
}
