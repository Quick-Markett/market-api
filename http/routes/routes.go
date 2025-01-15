package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	healthRoutes "main.go/http/routes/health"
	marketRoutes "main.go/http/routes/markets"
	ordersRoutes "main.go/http/routes/orders"
	productsRoutes "main.go/http/routes/products"
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

	healthPaths := r.Group("/health"); {
		healthRoutes.RegisterHealthRoutes(healthPaths)
	}

	log.Fatal(r.Run(":8080"))
}
