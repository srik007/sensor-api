package aggregators

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
	"gorm.io/gorm"
)

type AggregatorQueryHandler struct {
	DataStore *gorm.DB
}

type GroupAggregatorResult struct {
	AverageTemperature  float64
	AverageTransparency float64
	GroupName           string
}

func NewAggregatorQueryHandler(dataStore *gorm.DB) *AggregatorQueryHandler {
	return &AggregatorQueryHandler{
		DataStore: dataStore,
	}

}

func (a *AggregatorQueryHandler) CollectGroupAggregatorsByName(groupName string) GroupAggregatorResult {
	var groupAggregatorResult GroupAggregatorResult
	a.DataStore.Model(&entity.SensorData{}).
		Joins("JOIN sensors ON sensor_data.sensor_id = sensors.id").
		Joins("JOIN sensor_groups ON sensors.name = sensor_groups.name").
		Where("sensor_groups.name = ?", groupName).
		Select("AVG(sensor_data.value) as average_temperature, AVG(sensor_data.transparency) as average_transparency").
		Scan(&groupAggregatorResult)
	groupAggregatorResult.GroupName = groupName
	return groupAggregatorResult
}

func (a *AggregatorQueryHandler) CollectGroupAggregators() []GroupAggregatorResult {
	var uniqueGroupNames []string
	var groupAggregatorResults []GroupAggregatorResult
	a.DataStore.Model(&entity.SensorGroup{}).
		Select("DISTINCT name").
		Pluck("name", &uniqueGroupNames)
	for _, groupName := range uniqueGroupNames {
		var result GroupAggregatorResult
		a.DataStore.Model(&entity.SensorData{}).
			Joins("JOIN sensors ON sensor_data.sensor_id = sensors.id").
			Joins("JOIN sensor_groups ON sensors.name = sensor_groups.name").
			Where("sensor_groups.name = ?", groupName).
			Select("AVG(sensor_data.value) as average_temperature, AVG(sensor_data.transparency) as average_transparency").
			Scan(&result)
		result.GroupName = groupName
		groupAggregatorResults = append(groupAggregatorResults, result)
	}
	return groupAggregatorResults
}

func (a *AggregatorQueryHandler) CollectTopNSpeciesUnderGroup(groupName string, top int) entity.Species {
	uniqueSpecies := a.CollectSpeciesUnderGroup(groupName)
	sort.Slice(uniqueSpecies, func(i, j int) bool {
		return uniqueSpecies[i].Count > uniqueSpecies[j].Count
	})
	if len(uniqueSpecies) > top {
		uniqueSpecies = uniqueSpecies[:top]
	}
	return uniqueSpecies
}

func (a *AggregatorQueryHandler) CollectSpeciesUnderGroup(groupName string) entity.Species {
	var species []string
	var uniqueSpecies entity.Species
	a.DataStore.Model(&entity.SensorData{}).
		Joins("JOIN sensors ON sensor_data.sensor_id = sensors.id").
		Joins("JOIN sensor_groups ON sensors.name = sensor_groups.name").
		Where("sensor_groups.name = ?", groupName).
		Select("DISTINCT sensor_data.species").
		Pluck("species", &species)
	fmt.Println(species)
	for _, specie := range species {
		var newSpecie entity.Species
		json.Unmarshal([]byte(specie), &newSpecie)
		uniqueSpecies = append(uniqueSpecies, newSpecie...)
	}
	return uniqueSpecies
}

func (a *AggregatorQueryHandler) CalculateMinTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax float64) valueObjects.Temparature {

	var minTemparature valueObjects.Temparature

	a.DataStore.Model(&entity.SensorData{}).
		Joins("JOIN sensors ON sensor_data.sensor_id = sensors.id").
		Select("MIN(sensor_data.value) as value").
		Where("sensors.x BETWEEN ? AND ? AND sensors.y BETWEEN ? AND ? AND sensors.z BETWEEN ? AND ?", xMin, xMax, yMin, yMax, zMin, zMax).
		Scan(&minTemparature)
	minTemparature.Scale = "Celsius"
	return minTemparature
}

func (a *AggregatorQueryHandler) CalculateMaxTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax float64) valueObjects.Temparature {

	var maxTemparature valueObjects.Temparature

	a.DataStore.Model(&entity.SensorData{}).
		Joins("JOIN sensors ON sensor_data.sensor_id = sensors.id").
		Select("MAX(sensor_data.value) as value").
		Where("sensors.x BETWEEN ? AND ? AND sensors.y BETWEEN ? AND ? AND sensors.z BETWEEN ? AND ?", xMin, xMax, yMin, yMax, zMin, zMax).
		Scan(&maxTemparature)
	maxTemparature.Scale = "Celsius"
	return maxTemparature
}

func (a *AggregatorQueryHandler) CalculateAverageTemparatureBySensor(codeName entity.CodeName, startTime, endTime time.Time) valueObjects.Temparature {
	var sensorId uint
	var blah string
	var avgTemparature valueObjects.Temparature
	a.DataStore.Model(&entity.Sensor{}).
		Select("id").
		Where("sensors.name = ? AND sensors.group_id = ?", codeName.Name, codeName.GroupId).
		Scan(&sensorId)
	a.DataStore.Model(&entity.SensorData{}).
		Select("AVG(sensor_data.value)").
		Where("sensor_id = ? AND created_at BETWEEN ? AND ?", sensorId, startTime, endTime).
		Scan(&blah)
	avgTemparature.Scale = "Celsius"
	return avgTemparature
}
