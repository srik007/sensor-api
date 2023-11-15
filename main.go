package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/srik007/sensor-api/infrastructure/cache"
	"github.com/srik007/sensor-api/infrastructure/persistance"
	"github.com/srik007/sensor-api/interfaces"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/srik007/sensor-api/docs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

// @title Underwater Sensor API
// @version 1.0
// @description A Under water sensor api to monitor the sensor data and generate aggrates of the data.
// @host localhost:8080
// @BasePath /api/v1
func main() {

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	cacheUrl := os.Getenv("REDIS_URL")

	services, err := persistance.NewRepositories(dbdriver, user, password, port, host, dbname)
	cahce := cache.NewCache(cacheUrl)

	if err != nil {
		panic(err)
	}

	defer services.Close()

	services.Automigrate()

	sensorHandler := interfaces.NewSensorHandler(services.SensorRepository, services.SensorGroupRepository, services.DataStore, cahce)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{

		v1.POST("create-metadata", sensorHandler.CreateMetadata)

		v1.POST("schedule", sensorHandler.ScheduleJob)

		v1.GET("/group/:groupName/transparency", sensorHandler.CollectAverageTransparencyUnderGroup)

		v1.GET("/group/:groupName/temparature", sensorHandler.CollectAverageTemparatureUnderGroup)

		v1.GET("/group/:groupName/species", sensorHandler.CollectSpeciesUnderGroup)

		v1.GET("/group/:groupName/species/top/:topN", sensorHandler.CollectTopNSpeciesUnderGroup)

		v1.GET("/region/temperature/min", sensorHandler.CalculateMinTemparatureInsideARegion)

		v1.GET("/region/temperature/max", sensorHandler.CalculateMaxTemparatureInsideARegion)

	}

	app_port := os.Getenv("PORT")

	r.Run(":" + app_port)
}
