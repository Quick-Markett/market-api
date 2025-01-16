package categoriesRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/categories"
	middleware "main.go/middlewares"
)

func RegisterCategoriesRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateCategory)
	r.GET(":id", controllers.GetCategoryById)
	r.PUT(":id", controllers.UpdateCategory)
	r.DELETE(":id", controllers.DeleteCategory)
	r.GET("/get-market-categories/:id", controllers.GetMarketCategories)
}
