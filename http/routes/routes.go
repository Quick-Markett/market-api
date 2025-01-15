package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	ordersRoutes "main.go/http/routes/orders"
)

func HandleRequest(r *gin.Engine) {
	ordersPath := r.Group("/orders"); {
		ordersRoutes.RegisterOrdersRoutes(ordersPath)
	}

	log.Fatal(r.Run(":8080"))
}