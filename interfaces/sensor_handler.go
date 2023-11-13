package interfaces

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srik007/sensor-api/application"
	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
)

var sensorGroupData = []string{"alpha", "gamma", "beta"}

type Sensor struct {
	sensorApp      application.SensorAppInterface
	sensorGroupApp application.SensorGroupAppInterface
}

func NewSensor(sApp application.SensorAppInterface, sgApp application.SensorGroupAppInterface) *Sensor {
	return &Sensor{
		sensorApp:      sApp,
		sensorGroupApp: sgApp,
	}
}

func (s *Sensor) SaveSensorGroupData(c *gin.Context) {

	var sensorGroups []entity.SensorGroup

	for _, value := range sensorGroupData {
		sensorGroup := entity.SensorGroup{Name: value, SensorCount: 0}
		sensorGroups = append(sensorGroups, sensorGroup)
	}

	data, error := s.sensorGroupApp.SaveAll(sensorGroups)

	if error != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error.")
	}
	c.JSON(http.StatusCreated, data)
}

func (s *Sensor) SaveSensorData(c *gin.Context) {

	var sensors []entity.Sensor

	numberOfSensors := 10

	ocean3D := valueObjects.NewOcean3D(-90.0, 90.0, -180.0, 180.0, 0.0, 100.0)

	for i := 0; i < numberOfSensors; i++ {

		randomSensorGroupName := sensorGroupData[rand.Intn(len(sensorGroupData))]

		//Index should be fetch from entity layer and increment
		codeName := entity.CodeName{Name: randomSensorGroupName}

		coordiante := ocean3D.GetRandomCoordinates3D()

		fmt.Printf("%v", coordiante)

		dataOutputRate := valueObjects.DataOuputRate{Value: 10, Format: "seconds"}

		sensor := entity.Sensor{CodeName: codeName, Coordinate: coordiante, DataOutputRate: dataOutputRate}

		sensors = append(sensors, sensor)
	}

	data, error := s.sensorApp.SaveAll(sensors)

	if error != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error.")
	}
	c.JSON(http.StatusCreated, data)
}
