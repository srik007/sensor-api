package sensormetadata

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/repository"
	"github.com/srik007/sensor-api/domain/valueObjects"
)

type MetadataCreator struct {
	SensorRepository      repository.SensorRepository
	SensorGroupRepository repository.SensorGroupRepository
}

func NewMetadataCreator(s repository.SensorRepository, sg repository.SensorGroupRepository) *MetadataCreator {
	return &MetadataCreator{
		SensorRepository:      s,
		SensorGroupRepository: sg,
	}
}

var sensorGroupNames = []string{
	"Alpha", "Beta", "Gamma", "Delta", "Epsilon", "Zeta", "Eta", "Theta", "Iota", "Kappa",
	"Lambda", "Mu", "Nu", "Xi", "Omicron", "Pi", "Rho", "Sigma", "Tau", "Upsilon",
	"Phi", "Chi", "Psi", "Omega",
}

func (m *MetadataCreator) CreateSensorMetadata() {
	var sensors []entity.Sensor
	numberOfSensors, err := strconv.Atoi(os.Getenv("NUMBER_OF_SENSORS"))
	if err != nil {
		fmt.Println("Invalid number of sensors configured.")
		numberOfSensors = 4
	}
	ocean3D := valueObjects.NewOcean3D(-90.0, 90.0, -180.0, 180.0, 0.0, 100.0)
	for i := 0; i < numberOfSensors; i++ {
		randomSensorGroupName := sensorGroupNames[rand.Intn(len(sensorGroupNames))]
		codeName := entity.CodeName{Name: randomSensorGroupName}
		coordiante := ocean3D.GetRandomCoordinates3D()
		dataOutputRate := valueObjects.DataOuputRate{Value: 10, Format: "seconds"}
		sensor := entity.Sensor{CodeName: codeName, Coordinate: coordiante, DataOutputRate: dataOutputRate}
		sensors = append(sensors, sensor)
	}
	_, error := m.SensorRepository.SaveAll(sensors)
	if error != nil {
		fmt.Errorf("Failed to generate sensor data.")
	}
}

func (m *MetadataCreator) CreateSensorGroupMetaData() {
	var sensorGroups []entity.SensorGroup
	for _, value := range sensorGroupNames {
		sensorGroup := entity.SensorGroup{Name: value, SensorCount: 0}
		sensorGroups = append(sensorGroups, sensorGroup)
	}
	_, error := m.SensorGroupRepository.SaveAll(sensorGroups)
	if error != nil {
		fmt.Errorf("Failed to generate sensor groups.")
	}
}
