
package authRoutes

import (
	"github.com/gin-gonic/gin"
	google "main.go/http/controllers/auth/google"
	sso "main.go/http/controllers/auth/sso"
	jwt "main.go/http/controllers/auth/refreshToken"
	middleware "main.go/middlewares"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	// SSO
	r.POST("/sso/create-user", sso.CreateUserWithSSO)
	r.POST("/sso/login-user", sso.LoginUserWithSso)

	// Google
	r.POST("/google/create-user", google.CreateUserWithGoogle)
	r.POST("/google/login-user", google.LoginUserWithGoogle)

	// JWT Token
	r.POST("/refresh-token", jwt.RefreshToken)
}
