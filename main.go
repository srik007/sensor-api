package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/srik007/sensor-api/infrastructure/cache"
	"github.com/srik007/sensor-api/infrastructure/persistance"
	"github.com/srik007/sensor-api/interfaces"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

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

	r.POST("/api/generate", sensorHandler.GenerateMetadata)

	r.POST("/api/schedule", sensorHandler.ScheduleJob)

	r.GET("/group/:groupName/transparency", sensorHandler.CollectAverageTransparencyUnderGroup)

	r.GET("/group/:groupName/temparature", sensorHandler.CollectAverageTemparatureUnderGroup)

	r.GET("/group/:groupName/species", sensorHandler.CollectSpeciesUnderGroup)

	r.GET("/group/:groupName/species/top/:topN", sensorHandler.CollectTopNSpeciesUnderGroup)

	app_port := os.Getenv("PORT")

	r.Run(":" + app_port)
}
