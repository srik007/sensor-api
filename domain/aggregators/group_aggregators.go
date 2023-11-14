package aggregators

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/srik007/sensor-api/domain/entity"
	"gorm.io/gorm"
)

type Aggregator struct {
	DataStore *gorm.DB
}

func NewAggregator(dataStore *gorm.DB) *Aggregator {
	return &Aggregator{
		DataStore: dataStore,
	}
}

func (a *Aggregator) Run() {
	go CollectGroupAggregators(a.DataStore)
}

func CollectGroupAggregators(dataStore *gorm.DB) {

	for {
		time.Sleep(time.Duration(10) * time.Second)
		var uniqueGroupNames []string

		dataStore.Model(&entity.SensorGroup{}).
			Select("DISTINCT name").
			Pluck("name", &uniqueGroupNames)
		for _, groupName := range uniqueGroupNames {
			var result struct {
				GroupName           string
				AverageTemperature  float64
				AverageTransparency float64
			}
			dataStore.Model(&entity.SensorData{}).
				Joins("JOIN sensors ON sensor_data.sensor_id = sensors.id").
				Joins("JOIN sensor_groups ON sensors.name = sensor_groups.name").
				Where("sensor_groups.name = ?", groupName).
				Select("AVG(sensor_data.value) as average_temperature, AVG(sensor_data.transparency) as average_transparency").
				Scan(&result)

			result.GroupName = groupName
		}
	}
}

func (a *Aggregator) CollectTopNSpeciesUnderGroup(groupName string, top int) entity.Species {
	uniqueSpecies := a.CollectSpeciesUnderGroup(groupName)
	sort.Slice(uniqueSpecies, func(i, j int) bool {
		return uniqueSpecies[i].Count > uniqueSpecies[j].Count
	})
	if len(uniqueSpecies) > top {
		uniqueSpecies = uniqueSpecies[:top]
	}
	return uniqueSpecies
}

func (a *Aggregator) CollectSpeciesUnderGroup(groupName string) entity.Species {
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
