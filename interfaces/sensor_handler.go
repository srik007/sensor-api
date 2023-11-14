package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/repository"
)

type SensorHandler struct {
	sensorApp application.SensorAppInterface
}

func NewSensorHandler(s repository.SensorRepository, sg repository.SensorGroupRepository) *SensorHandler {
	return &SensorHandler{
		sensorApp: &application.SensorApp{SensorRepository: s, SensorGroupRepository: sg},
	}
}

func (s *SensorHandler) GenerateMetadata(c *gin.Context) {
	s.sensorApp.GenerateMetadata()
	c.JSON(200, "Data generated successfully.")
}

func (s *SensorHandler) ScheduleJob(c *gin.Context) {
	s.sensorApp.ScheduleJob()
	c.JSON(202, "Triggred monitoring job.")
}
