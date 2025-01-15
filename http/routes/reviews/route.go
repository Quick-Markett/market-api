package reviewsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/reviews"
	middleware "main.go/middlewares"
)

func RegisterReviewRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateReview)
	r.GET(":id", controllers.GetReviewById)
	r.PUT(":id", controllers.UpdateReview)
	r.DELETE(":id", controllers.DeleteReview)
	r.GET("/get-order-review/:id", controllers.GetOrderReview)
}
