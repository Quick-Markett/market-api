package userRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/users"
	middleware "main.go/middlewares"
)

func RegisterUsersRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", controllers.GetUsers)
	r.GET(":id", controllers.GetUser)
	r.POST("", controllers.CreateUser)
	r.PUT(":id", controllers.UpdateUser)
	r.DELETE(":id", controllers.DeleteUser)

}
