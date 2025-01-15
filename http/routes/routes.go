package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	healthRoutes "main.go/http/routes/health"
	marketRoutes "main.go/http/routes/markets"
	orderItemsRoutes "main.go/http/routes/orderItems"
	ordersRoutes "main.go/http/routes/orders"
	productsRoutes "main.go/http/routes/products"
	reviewsRoutes "main.go/http/routes/reviews"
	userRoutes "main.go/http/routes/users"
)

func HandleRequest(r *gin.Engine) {
	ordersPath := r.Group("/orders"); {
		ordersRoutes.RegisterOrdersRoutes(ordersPath)
	}

	usersPath := r.Group("/users"); {
		userRoutes.RegisterUsersRoutes(usersPath)
	}

	marketsPath := r.Group("/markets"); {
		marketRoutes.RegisterMarketsRoutes(marketsPath)
	}

	productsPath := r.Group("/products"); {
		productsRoutes.RegisterProductsRoutes(productsPath)
	}

	orderItemsPath := r.Group("/order-items"); {
		orderItemsRoutes.RegisterOrderItemsRoutes(orderItemsPath)
	}

	reviewsPath := r.Group("/reviews"); {
		reviewsRoutes.RegisterReviewRoutes(reviewsPath)
	}

	healthPaths := r.Group("/health"); {
		healthRoutes.RegisterHealthRoutes(healthPaths)
	}

	log.Fatal(r.Run(":8080"))
}
