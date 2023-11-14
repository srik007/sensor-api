package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/repository"
	"gorm.io/gorm"
)

type SensorHandler struct {
	sensorApp application.SensorAppInterface
}

func NewSensorHandler(s repository.SensorRepository, sg repository.SensorGroupRepository, dataStore *gorm.DB) *SensorHandler {
	return &SensorHandler{
		sensorApp: application.NewSensorApp(s, sg, dataStore),
	}
}

func (s *SensorHandler) GenerateMetadata(c *gin.Context) {
	s.sensorApp.GenerateMetadata()
	c.JSON(http.StatusOK, "Data generated successfully.")
}

func (s *SensorHandler) ScheduleJob(c *gin.Context) {
	s.sensorApp.ScheduleJob()
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
