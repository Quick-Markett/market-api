package orderItemsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/orderItems"
	middleware "main.go/middlewares"
)

func RegisterOrderItemsRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", middleware.JWTMiddleware(), controllers.CreateOrderItem)
	r.GET(":id", middleware.JWTMiddleware(), controllers.GetOrderItemById)
	r.PUT(":id", middleware.JWTMiddleware(), controllers.UpdateOrderItem)
	r.DELETE(":id", middleware.JWTMiddleware(), controllers.DeleteOrderItem)
	r.GET("/get-order-items/:id", middleware.JWTMiddleware(), controllers.GetOrderItems)
}
