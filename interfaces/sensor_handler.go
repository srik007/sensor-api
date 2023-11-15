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

type AverageTransparencyResponse struct {
	Transparency float64
	GroupName    string
}

type AverageTemparatureResponse struct {
	Temparature float64
	GroupName   string
}

func NewSensorHandler(s repository.SensorRepository, sg repository.SensorGroupRepository, dataStore *gorm.DB, cache *cache.RedisCache) *SensorHandler {
	return &SensorHandler{
		sensorApp:  application.NewSensorApp(s, sg, dataStore, cache),
		CacheStore: cache,
	}
}

// CreateMetadata is a handler that creates metada for sensors & sensors groups
// @Summary Create sensor metadata and sensor group metadata
// @Description Create the meta data for sensors and sensor groups
// @ID create-sensor-metadata
// @Success 200 {text} Sensor metadata created successfully.
// @Router /createMetadata [post]
func (s *SensorHandler) CreateMetadata(c *gin.Context) {
	s.sensorApp.CreateMetadata()
	c.JSON(http.StatusOK, "Sensor metadata created successfully.")
}

// @Summary Background job to sendule sensors and generate the data for given sensors & Also schedule the backgorund job to do the aggregations on generated data.
// @Description Run background job to generate sensors data
// @ID schedule-jobs
// @Success 202 {text} Triggered monitoring job..
// @Router /schedule [post]
func (s *SensorHandler) ScheduleJob(c *gin.Context) {
	s.sensorApp.Schedule(c)
	c.JSON(http.StatusAccepted, "Triggred monitoring job.")
}

// CollectSpeciesUnderGroup is a handler that retrieves species details by group name.
// @Summary Collect all the fish species detected under sensors belonging to a given sensor group.
// @Description Get total species by group name.
// @ID collect-total-species-under-group
// @Success 200 {object} entity.Species "Successful response"
// @Produce json
// @Param groupName path string true "Group name"
// @Router /group/{groupName}/species [get]
func (s *SensorHandler) CollectSpeciesUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	species := s.sensorApp.CollectSpeciesUnderGroup(groupName)
	c.JSON(http.StatusOK, species)
}

// CollectTopNSpeciesUnderGroup is a handler that retrieves top N species details based on count by group name.
// @Summary Collect top N fish species detected under sensors belonging to a given sensor group.
// @Description Get top N species by group name.
// @ID collect-top-n-species-under-group
// @Success 200 {object} entity.Species "Successful response"
// @Produce json
// @Param groupName path string true "Group name"
// @Param topN path int true "Top n"
// @Router /group/{groupName}/species//top/:topN [get]
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

// CollectAverageTransparencyUnderGroup is a handler that retrives average tranpercy of all the sensors under given group.
// @Summary Collect average transparency of all the sensors transparency under given group
// @Description Collect average transparency
// @ID collect-avg-transparency
// @Success 200 {object} AverageTransparencyResponse "Successful response"
// @Produce json
// @Param groupName path string true "Group name"
// @Router /group/{groupName}/temparature [get]
func (s *SensorHandler) CollectAverageTransparencyUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	result := s.CacheStore.VerifyCache(c, groupName)
	if result != nil {
		cahcedResponse := &AverageTransparencyResponse{Transparency: result.AverageTransparency, GroupName: groupName}
		c.JSON(http.StatusOK, cahcedResponse)
		return
	}
	avgTransparency := s.sensorApp.CollectAverageTransparencyUnderGroup(groupName)
	averageTransparencyResponse := &AverageTransparencyResponse{
		Transparency: avgTransparency,
		GroupName:    groupName,
	}
	c.JSON(http.StatusOK, averageTransparencyResponse)
	return
}

// CollectAverageTemparatureUnderGroup is a handler that retrives average temparature of all the sensors under given group.
// @Summary Collect average temparature of all the sensors temparature under given group
// @Description Collect average temparature
// @ID collect-avg-temparature
// @Success 200 {object} AverageTransparencyResponse "Successful response"
// @Produce json
// @Param groupName path string true "Group name"
// @Router /group/{groupName}/temparature [get]
func (s *SensorHandler) CollectAverageTemparatureUnderGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	result := s.CacheStore.VerifyCache(c, groupName)
	if result != nil {
		cahcedResponse := &AverageTemparatureResponse{Temparature: result.AverageTemperature, GroupName: groupName}
		c.JSON(http.StatusOK, cahcedResponse)
		return
	}
	avgTemparature := s.sensorApp.CollectAverageTemparatureUnderGroup(groupName)
	averageTemparatureResponse := &AverageTemparatureResponse{Temparature: avgTemparature, GroupName: groupName}
	c.JSON(http.StatusOK, averageTemparatureResponse)
	return
}

type ErrorResposne struct {
	Message string
}

// CalculateMinTemparatureInsideARegion is a handler that calculates min temparature in a given region.
// @Summary Calculate minimum temparature inside a region
// @Description Calculate minimum temparature
// @ID calculate-min-temparature
// @Success 200 {object} valueObjects.Temparature "Successful response"
// @Failure 400 {object} ErrorResposne "Failure response"
// @Produce json
// @Param xMin query float64 true "Minimum x"
// @Param xMax query float64 true "Maximum x"
// @Param yMin query float64 true "Minimum y"
// @Param yMax query float64 true "Maximum y"
// @Param zMin query float64 true "Minimum z"
// @Param zMax query float64 true "Maximum z"
// @Router /region/temperature/min [get]
func (s *SensorHandler) CalculateMinTemparatureInsideARegion(c *gin.Context) {
	xMin, _ := strconv.ParseFloat(c.Query("xMin"), 64)
	xMax, _ := strconv.ParseFloat(c.Query("xMax"), 64)
	yMin, _ := strconv.ParseFloat(c.Query("yMin"), 64)
	yMax, _ := strconv.ParseFloat(c.Query("yMax"), 64)
	zMin, _ := strconv.ParseFloat(c.Query("zMin"), 64)
	zMax, _ := strconv.ParseFloat(c.Query("zMax"), 64)
	if xMin >= xMax || yMin >= yMax || zMin >= zMax {
		errorResponse := &ErrorResposne{Message: "Invalid range parameters"}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	minTemperature := s.sensorApp.CalculateMinTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax)
	c.JSON(http.StatusOK, minTemperature)
}

// CalculateMaxTemparatureInsideARegion is a handler that calculates max temparature in a given region.
// @Summary Calculate maximum temparature inside a region
// @Description Calculate maximum temparature
// @ID calculate-max-temparature
// @Success 200 {object} valueObjects.Temparature "Successful response"
// @Failure 400 {object} ErrorResposne "Failure response"
// @Produce json
// @Param xMin query float64 true "Minimum x"
// @Param xMax query float64 true "Maximum x"
// @Param yMin query float64 true "Minimum y"
// @Param yMax query float64 true "Maximum y"
// @Param zMin query float64 true "Minimum z"
// @Param zMax query float64 true "Maximum z"
// @Router /region/temperature/max [get]
func (s *SensorHandler) CalculateMaxTemparatureInsideARegion(c *gin.Context) {
	xMin, _ := strconv.ParseFloat(c.Query("xMin"), 64)
	xMax, _ := strconv.ParseFloat(c.Query("xMax"), 64)
	yMin, _ := strconv.ParseFloat(c.Query("yMin"), 64)
	yMax, _ := strconv.ParseFloat(c.Query("yMax"), 64)
	zMin, _ := strconv.ParseFloat(c.Query("zMin"), 64)
	zMax, _ := strconv.ParseFloat(c.Query("zMax"), 64)
	if xMin >= xMax || yMin >= yMax || zMin >= zMax {
		errorResponse := &ErrorResposne{Message: "Invalid range parameters"}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	maxTemparature := s.sensorApp.CalculateMaxTemparatureInsideARegion(xMin, xMax, yMin, yMax, zMin, zMax)
	c.JSON(http.StatusOK, maxTemparature)
}
