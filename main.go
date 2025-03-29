package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/http/routes"
	middlewares "main.go/middlewares"
)

// var ginLambda *ginadapter.GinLambda

func main() {
	r := gin.Default()

	middlewares.PrometheusInit()
	r.Use(middlewares.TrackMetrics())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	database.ConnectWithDatabase()
	routes.HandleRequest(r)
	// ginLambda = ginadapter.New(r)
}

// func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return ginLambda.Proxy(req)
// }
