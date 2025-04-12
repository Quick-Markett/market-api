package marketsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/markets"
	middleware "main.go/middlewares"
)

func RegisterMarketsRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", controllers.GetMarkets)
	r.GET(":id", controllers.GetMarket)
	r.GET("/get-market-by-slug/:slug", controllers.GetMarketBySlug)
	r.POST("", middleware.JWTMiddleware(), controllers.CreateMarket)
	r.PUT(":id", middleware.JWTMiddleware(), controllers.UpdateMarket)
	r.DELETE(":id", middleware.JWTMiddleware(), controllers.DeleteMarket)

	// NEARBY MARKETS
	r.POST("/nearby", middleware.JWTMiddleware(), controllers.FindMarketsWithinRadius)
}
