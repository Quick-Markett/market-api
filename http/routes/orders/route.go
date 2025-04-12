package ordersRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/orders"
	middleware "main.go/middlewares"
)

func RegisterOrdersRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", middleware.JWTMiddleware(), controllers.CreateOrder)
	r.GET(":id", middleware.JWTMiddleware(), controllers.GetOrderById)
	r.PUT(":id", middleware.JWTMiddleware(), controllers.UpdateOrder)
	r.DELETE(":id", middleware.JWTMiddleware(), controllers.DeleteOrder)
	r.GET("/get-market-orders/:id", middleware.JWTMiddleware(), controllers.GetMarketOrders)
	r.GET("/get-user-orders/:userId", middleware.JWTMiddleware(), controllers.GetUserOrderItems)
}
