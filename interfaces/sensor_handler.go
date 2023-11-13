package interfaces

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
	"github.com/srik007/sensor-api/infrastructure/monitor"
)

type SensorHandler struct {
	sensorApp      application.SensorAppInterface
	sensorGroupApp application.SensorGroupAppInterface
}

var sensorGroupNames = []string{
	"Alpha", "Beta", "Gamma", "Delta", "Epsilon", "Zeta", "Eta", "Theta", "Iota", "Kappa",
	"Lambda", "Mu", "Nu", "Xi", "Omicron", "Pi", "Rho", "Sigma", "Tau", "Upsilon",
	"Phi", "Chi", "Psi", "Omega",
}

func NewSensorHandler(sApp application.SensorAppInterface, sgApp application.SensorGroupAppInterface) *SensorHandler {
	return &SensorHandler{
		sensorApp:      sApp,
		sensorGroupApp: sgApp,
	}
}

func (s *SensorHandler) Generate(c *gin.Context) {
	var sensorGroups []entity.SensorGroup
	for _, value := range sensorGroupNames {
		sensorGroup := entity.SensorGroup{Name: value, SensorCount: 0}
		sensorGroups = append(sensorGroups, sensorGroup)
	}
	_, error := s.sensorGroupApp.SaveAll(sensorGroups)
	if error != nil {
		fmt.Errorf("Failed to generate sensor groups.")
	}

	var sensors []entity.Sensor
	numberOfSensors := 10
	ocean3D := valueObjects.NewOcean3D(-90.0, 90.0, -180.0, 180.0, 0.0, 100.0)
	for i := 0; i < numberOfSensors; i++ {
		randomSensorGroupName := sensorGroupNames[rand.Intn(len(sensorGroupNames))]
		codeName := entity.CodeName{Name: randomSensorGroupName}
		coordiante := ocean3D.GetRandomCoordinates3D()
		dataOutputRate := valueObjects.DataOuputRate{Value: 10, Format: "seconds"}
		sensor := entity.Sensor{CodeName: codeName, Coordinate: coordiante, DataOutputRate: dataOutputRate}
		sensors = append(sensors, sensor)
	}
	_, error = s.sensorApp.SaveAll(sensors)
	if error != nil {
		fmt.Errorf("Failed to generate sensors")
	}

	c.JSON(200, "Data generated successfully.")
}

func (s *SensorHandler) Monitor(c *gin.Context) {
	sensors := s.sensorApp.GetAll()
	if len(sensors) > 0 {
		monitor.Register(sensors)
	}
	c.JSON(202, "Triggred monitoring job.")
}
