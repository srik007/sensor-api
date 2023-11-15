package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/repository"
	"github.com/srik007/sensor-api/infrastructure/cache"
	"gorm.io/gorm"
)

type SensorHandler struct {
	sensorApp  application.SensorAppInterface
	CacheStore *cache.RedisCache
}

func NewSensorHandler(s repository.SensorRepository, sg repository.SensorGroupRepository, dataStore *gorm.DB, cache *cache.RedisCache) *SensorHandler {
	return &SensorHandler{
		sensorApp:  application.NewSensorApp(s, sg, dataStore, cache),
		CacheStore: cache,
	}
}

func (s *SensorHandler) GenerateMetadata(c *gin.Context) {
	s.sensorApp.GenerateMetadata()
	c.JSON(http.StatusOK, "Data generated successfully.")
}

func (s *SensorHandler) ScheduleJob(c *gin.Context) {
	s.sensorApp.Schedule(c)
	c.JSON(http.StatusAccepted, "Triggred monitoring job.")
}

func (s *SensorHandler) CollectSpeciesUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	species := s.sensorApp.CollectSpeciesUnderGroup(groupName)
	c.JSON(http.StatusOK, species)
}

func (s *SensorHandler) CollectTopNSpeciesUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	topNValue := c.Param("topN")
	topN, err := strconv.Atoi(topNValue)
	if err != nil {
		fmt.Println("Invalid number:", err)
		c.JSON(http.StatusBadRequest, "Wrong topN number")
	}
	species := s.sensorApp.CollectTopNSpeciesUnderGroup(groupName, topN)
	c.JSON(http.StatusOK, species)
}

func (s *SensorHandler) CollectAverageTransparencyUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	result := s.CacheStore.VerifyCache(c, groupName)
	if result != nil {
		c.JSON(http.StatusOK, gin.H{"Transparency": result.AverageTransparency, "GroupName": groupName})
	}
	avgTransparency := s.sensorApp.CollectAverageTransparencyUnderGroup(groupName)
	c.JSON(http.StatusOK, gin.H{"Transparency": avgTransparency, "GroupName": groupName})
	return
}

func (s *SensorHandler) CollectAverageTemparatureUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	result := s.CacheStore.VerifyCache(c, groupName)
	if result != nil {
		c.JSON(http.StatusOK, gin.H{"Temparature": result.AverageTransparency, "GroupName": groupName})
		return
	}
	avgTemparature := s.sensorApp.CollectAverageTemparatureUnderGroup(groupName)
	c.JSON(http.StatusOK, gin.H{"Temparature": avgTemparature, "GroupName": groupName})
}
