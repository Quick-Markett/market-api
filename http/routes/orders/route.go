package ordersRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/orders"
	middleware "main.go/middlewares"
)

func RegisterOrdersRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateOrder)
}
