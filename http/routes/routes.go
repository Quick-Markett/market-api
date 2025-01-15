package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	marketRoutes "main.go/http/routes/markets"
	ordersRoutes "main.go/http/routes/orders"
	userRoutes "main.go/http/routes/users"
)

func HandleRequest(r *gin.Engine) {
	ordersPath := r.Group("/orders")
	{
		ordersRoutes.RegisterOrdersRoutes(ordersPath)
	}

	usersPath := r.Group("/users")
	{
		userRoutes.RegisterUsersRoutes(usersPath)
	}

	marketsPath := r.Group("/markets")
	{
		marketRoutes.RegisterMarketsRoutes(marketsPath)
	}

	log.Fatal(r.Run(":8080"))
}
