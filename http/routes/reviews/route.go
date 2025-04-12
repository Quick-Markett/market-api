package reviewsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/reviews"
	middleware "main.go/middlewares"
)

func RegisterReviewRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", middleware.JWTMiddleware(), controllers.CreateReview)
	r.GET(":id", controllers.GetReviewById)
	r.PUT(":id", middleware.JWTMiddleware(), controllers.UpdateReview)
	r.DELETE(":id", middleware.JWTMiddleware(), controllers.DeleteReview)
	r.GET("/get-order-review/:id", controllers.GetOrderReview)
}
