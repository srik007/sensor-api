package generators

import (
	"fmt"
	"math/rand"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	"github.com/srik007/sensor-api/domain/valueObjects"
)

type Generator struct {
	SensorRepository      repository.SensorRepository
	SensorGroupRepository repository.SensorGroupRepository
}

func NewGenerator(s repository.SensorRepository, sg repository.SensorGroupRepository) *Generator {
	return &Generator{
		SensorRepository:      s,
		SensorGroupRepository: sg,
	}
}

var sensorGroupNames = []string{
	"Alpha", "Beta", "Gamma", "Delta", "Epsilon", "Zeta", "Eta", "Theta", "Iota", "Kappa",
	"Lambda", "Mu", "Nu", "Xi", "Omicron", "Pi", "Rho", "Sigma", "Tau", "Upsilon",
	"Phi", "Chi", "Psi", "Omega",
}

func (g *Generator) GenerateSensorMetaData() {
	var sensors []entity.Sensor
	numberOfSensors := 1
	ocean3D := valueObjects.NewOcean3D(-90.0, 90.0, -180.0, 180.0, 0.0, 100.0)
	for i := 0; i < numberOfSensors; i++ {
		randomSensorGroupName := sensorGroupNames[rand.Intn(len(sensorGroupNames))]
		codeName := entity.CodeName{Name: randomSensorGroupName}
		coordiante := ocean3D.GetRandomCoordinates3D()
		dataOutputRate := valueObjects.DataOuputRate{Value: 10, Format: "seconds"}
		sensor := entity.Sensor{CodeName: codeName, Coordinate: coordiante, DataOutputRate: dataOutputRate}
		sensors = append(sensors, sensor)
	}
	_, error := g.SensorRepository.SaveAll(sensors)
	if error != nil {
		fmt.Errorf("Failed to generate sensor data.")
	}
}

func (g *Generator) GenerateSensorGroupMetaData() {
	var sensorGroups []entity.SensorGroup
	for _, value := range sensorGroupNames {
		sensorGroup := entity.SensorGroup{Name: value, SensorCount: 0}
		sensorGroups = append(sensorGroups, sensorGroup)
	}
	_, error := g.SensorGroupRepository.SaveAll(sensorGroups)
	if error != nil {
		fmt.Errorf("Failed to generate sensor groups.")
	}

}
