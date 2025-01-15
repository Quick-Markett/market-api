package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	marketGroup := r.Group("/markets")
	{
		marketGroup.GET("/", controllers.GetMarkets)
		marketGroup.GET("/:id", controllers.GetMarket)
		marketGroup.POST("/", controllers.CreateMarket)
		marketGroup.PUT("/:id", controllers.UpdateMarket)
		marketGroup.DELETE("/:id", controllers.DeleteMarket)
	}

	return r
}
