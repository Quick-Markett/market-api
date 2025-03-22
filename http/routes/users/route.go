package userRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/users"
	google "main.go/http/controllers/users/google"
	sso "main.go/http/controllers/users/sso"
	middleware "main.go/middlewares"
)

func RegisterUsersRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.GET("", controllers.GetUsers)
	r.GET(":id", controllers.GetUser)
	r.PUT(":id", controllers.UpdateUser)
	r.DELETE(":id", controllers.DeleteUser)

	// SSO
	r.POST("/sso/create-user", sso.CreateUserWithSSO)
	r.GET("/sso/login-user/:email", sso.LoginUserWithSso)

	// Google
	r.POST("/google/create-user", google.LoginUserWithGoogle)
	r.GET("/google/login-user/:googleId", google.LoginUserWithGoogle)
}
