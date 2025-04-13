package categoriesRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/categories"
	middleware "main.go/middlewares"
)

func RegisterCategoriesRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", middleware.JWTMiddleware(), controllers.CreateCategory)
	r.GET(":id", middleware.JWTMiddleware(), controllers.GetCategoryById)
	r.PUT(":id", middleware.JWTMiddleware(), controllers.UpdateCategory)
	r.DELETE(":id", middleware.JWTMiddleware(), controllers.DeleteCategory)
	r.GET("/get-market-categories", middleware.JWTMiddleware(), controllers.GetMarketCategories)
}
