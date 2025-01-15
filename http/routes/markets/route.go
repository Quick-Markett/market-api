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
	r.POST("", controllers.CreateMarket)
	r.PUT(":id", controllers.UpdateMarket)
	r.DELETE(":id", controllers.DeleteMarket)

}
