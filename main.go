package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/http/routes"
	middleware "main.go/middlewares"
)

func main() {
	r := gin.Default()

	middleware.PrometheusInit()
	r.Use(middleware.TrackMetrics())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	database.ConnectWithDatabase()
	routes.HandleRequest(r)

	fmt.Println("Iniciando projeto do mercadinho...")
}
