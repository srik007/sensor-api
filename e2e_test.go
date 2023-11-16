package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/srik007/sensor-api/domain/entity"
	"github.com/srik007/sensor-api/domain/valueObjects"
	"github.com/srik007/sensor-api/interfaces"
	"github.com/stretchr/testify/assert"
)

const BASE_URL = "http://localhost:8080/api/v1"

func TestAverageTrasnparencyInTheGroup(t *testing.T) {
	var averageTransparency interfaces.AverageTransparencyResponse
	response, err := http.Get(BASE_URL + "/group/Group1/transparency")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &averageTransparency)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, averageTransparency.GroupName, "Group1")
	assert.Equal(t, averageTransparency.Transparency, float64(62.5))
}

func TestAverageTemparatureInTheGroup(t *testing.T) {
	var averageTemparature interfaces.AverageTemparatureResponse
	response, err := http.Get(BASE_URL + "/group/Group2/temparature")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &averageTemparature)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, averageTemparature.Temparature, float64(250))
}

func TestTotalSpeciesInTheGroup(t *testing.T) {
	var species entity.Species
	response, err := http.Get(BASE_URL + "/group/Group2/species")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &species)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, len(species), 4)
}

func TestTopNSpeciesInTheGroup(t *testing.T) {
	var species entity.Species
	response, err := http.Get(BASE_URL + "/group/Group2/species/top/2")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &species)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, len(species), 2)
	assert.Equal(t, 10, species[0].Count)
}

func TestTopNSpeciesInTheGroupBetween(t *testing.T) {
	var species entity.Species
	response, err := http.Get(BASE_URL + "/group/Group2/species/top/2?from=1700137279&till=1700137279")
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &species)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, len(species), 0)
}

func TestMaxTemparatureInsideAGroup(t *testing.T) {
	var maxTemparature valueObjects.Temparature
	response, err := http.Get(BASE_URL + "/region/temparature/max?xMin=0&xMax=20&yMin=0&yMax=25&zMin=0&zMax=150")
	body, err := io.ReadAll(response.Body)
	fmt.Println(body)
	err = json.Unmarshal(body, &maxTemparature)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()
	fmt.Println(maxTemparature)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(300), maxTemparature.Value)
	assert.Equal(t, "Celsius", maxTemparature.Scale)
}

func TestMinTemparatureInsideAGroup(t *testing.T) {
	var minTemperature valueObjects.Temparature
	response, err := http.Get(BASE_URL + "/region/temparature/min?xMin=0&xMax=20&yMin=0&yMax=25&zMin=0&zMax=150")
	body, err := io.ReadAll(response.Body)
	fmt.Println(body)
	err = json.Unmarshal(body, &minTemperature)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(200), minTemperature.Value)
	assert.Equal(t, "Celsius", minTemperature.Scale)
}

func TestAverageTemparatureDetectedBySensor(t *testing.T) {
	var averageTemparature valueObjects.Temparature
	currentTimestamp := int64(time.Now().Unix())
	response, err := http.Get(BASE_URL + "/sensor/Group1_1/temparature/average?from=1697438904&till=" + strconv.FormatInt(currentTimestamp, 10))
	body, err := io.ReadAll(response.Body)
	fmt.Println(body)
	err = json.Unmarshal(body, &averageTemparature)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(200), averageTemparature.Value)
	assert.Equal(t, "Celsius", averageTemparature.Scale)
}
