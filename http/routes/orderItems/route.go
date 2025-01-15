package orderItemsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/orderItems"
	middleware "main.go/middlewares"
)

func RegisterOrderItemsRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateOrderItem)
	r.GET(":id", controllers.GetOrderItemById)
	r.PUT(":id", controllers.UpdateOrderItem)
	r.DELETE(":id", controllers.DeleteOrderItem)
	r.GET("/get-order-items/:id", controllers.GetOrderItems)
}
