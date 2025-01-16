package productsRoutes

import (
	"github.com/gin-gonic/gin"
	controllers "main.go/http/controllers/products"
	middleware "main.go/middlewares"
)

func RegisterProductsRoutes(r *gin.RouterGroup) {
	r.Use(middleware.ContentTypeMiddleware())

	r.POST("", controllers.CreateProduct)
	r.GET(":id", controllers.GetProductById)
	r.PUT(":id", controllers.UpdateProduct)
	r.DELETE(":id", controllers.DeleteProduct)
	r.GET("/get-market-products/:id", controllers.GetMarketProducts)
	r.GET("/filter", controllers.GetFilteredProducts)
}
