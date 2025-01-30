package main

import (
	"fmt"

	"main.go/database"
	"main.go/http/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	middlewares "main.go/middlewares"
)

var ginLambda *ginadapter.GinLambda

func init() {
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
	ginLambda = ginadapter.New(r)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(req)
}

func main() {
	fmt.Println("Iniciando projeto do mercadinho...")
	lambda.Start(Handler)
}
