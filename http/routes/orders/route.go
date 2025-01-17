package ordersRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/orders"
	middleware "main.go/middlewares"
)

func RegisterOrdersRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateOrder)
	r.GET(":id", controllers.GetOrderById)
	r.PUT(":id", controllers.UpdateOrder)
	r.DELETE(":id", controllers.DeleteOrder)
	r.GET("/get-market-orders/:id", controllers.GetMarketOrders)
	r.GET("/get-user-orders/:id", controllers.GetUserOrderItems)
}
