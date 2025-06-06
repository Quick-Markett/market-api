package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	getParams "main.go/pkg/params"
	response "main.go/pkg/response"
)

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to create products.")
		return
	}

	result := database.PostgresInstance.Create(&newProduct)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new product.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newProduct, nil, "")
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	result := database.PostgresInstance.Preload("Market").First(&product, id)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Product not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, product, nil, "")
}

func GetMarketProducts(c *gin.Context) {
	slug := c.Param("slug")

	var market models.Market
	if err := database.PostgresInstance.Where("slug = ?", slug).First(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found.")
		return
	}

	category := c.Query("category")
	var products []models.Product

	query := database.PostgresInstance.Where("market_id = ?", market.ID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	result := query.Find(&products)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve products.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, products, nil, "")
}

func UpdateProduct(c *gin.Context) {
	productId := c.Param("id")

	var request struct {
		Name        string `json:"product_name"`
		Description string `json:"product_description"`
		Slug        string `json:"slug"`
		Image       string `json:"product_image"`
		UnitPrice   string `json:"unit_price"`
		Stock       string `json:"stock"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to update products.")
		return
	}

	var product models.Product

	if err := database.PostgresInstance.First(&product, productId).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Product not found.")
		return
	}

	result := database.PostgresInstance.Model(&product).Updates(map[string]interface{}{
		"product_name":        request.Name,
		"product_description": request.Description,
		"slug":                request.Slug,
		"product_image":       request.Image,
		"unit_price":          request.UnitPrice,
		"stock":               request.Stock,
	})

	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the category.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, result, nil, "")
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.PostgresInstance.First(&product, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Product not found.")
		return
	}

	if err := database.PostgresInstance.Delete(&product).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete the product.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, nil, nil, "")
}

func GetMappedProductsByCategory(c *gin.Context) {
	marketSlug := c.Param("slug")
	var market models.Market

	if err := database.PostgresInstance.Where("slug = ?", marketSlug).First(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found.")
		return
	}

	var categories []models.Category
	if err := database.PostgresInstance.
		Preload("Products").
		Where("market_id = ?", market.ID).
		Find(&categories).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Error fetching categories.")
		return
	}

	var mappedItems []map[string]interface{}
	for _, category := range categories {
		var products []map[string]interface{}
		for _, product := range category.Products {
			if product.IsActive {
				products = append(products, map[string]interface{}{
					"id":                  product.ID,
					"product_name":        product.ProductName,
					"slug":                product.Slug,
					"unit_price":          product.UnitPrice,
					"product_image":       product.ProductImage,
					"product_description": product.ProductDescription,
				})
			}
		}

		if len(products) > 0 {
			mappedItems = append(mappedItems, map[string]interface{}{
				"category": category.Name,
				"products": products,
			})
		}
	}

	response.SendGinResponse(c, http.StatusOK, mappedItems, nil, "")
}

func GetFilteredProducts(c *gin.Context) {
	var products []models.Product

	category, hasCategory := getParams.GetParams(c, "category")
	priceMin, hasPriceMin := getParams.GetParams(c, "price_min")
	priceMax, hasPriceMax := getParams.GetParams(c, "price_max")

	query := database.PostgresInstance

	if hasCategory {
		query = query.Where("category = ?", category)
	}
	if hasPriceMin {
		query = query.Where("price >= ?", priceMin)
	}
	if hasPriceMax {
		query = query.Where("price <= ?", priceMax)
	}

	result := query.Find(&products)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve products.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, products, nil, "")
}
