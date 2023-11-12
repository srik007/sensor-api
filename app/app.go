package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	database *sql.DB
	Routes   *gin.Engine
}

func (app *App) createConnection() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, databaseName)
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	app.database = database
}

func (app *App) createRoutes() {
	routes := gin.Default()
	routes.GET("/hello", func(g *gin.Context) { g.JSON(200, nil) })
}

func (app *App) run() {
	app.Routes.Run(":8080")
}
