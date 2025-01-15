package healthRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/health"
	middleware "main.go/middlewares"
)

func RegisterHealthRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", controllers.Health)
}
