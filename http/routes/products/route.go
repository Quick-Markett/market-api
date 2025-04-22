package productsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/products"
	middleware "main.go/middlewares"
)

func RegisterProductsRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", middleware.JWTMiddleware(), controllers.CreateProduct)
	r.GET(":id", controllers.GetProductById)
	r.PUT("/:id", middleware.JWTMiddleware(), controllers.UpdateProduct)
	r.DELETE(":id", middleware.JWTMiddleware(), controllers.DeleteProduct)
	r.GET("/get-market-products/:slug", controllers.GetMarketProducts)
	r.GET("/filter", controllers.GetFilteredProducts)
	r.GET("/get-mapped-products/:slug", controllers.GetMappedProductsByCategory)
}
