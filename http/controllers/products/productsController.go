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
	marketID := c.Param("id")
	category := c.Query("category")
	var products []models.Product

	query := database.PostgresInstance.Where("market_id = ?", marketID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	result := query.Find(&products)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve products.")
		return
	}

	if len(products) == 0 {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "No products found for this market.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, products, nil, "")
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.PostgresInstance.First(&product, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Product not found.")
		return
	}

	var updatedData models.Product
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON for updating the product.")
		return
	}

	result := database.PostgresInstance.Model(&product).Updates(updatedData)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the product.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, product, nil, "")
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

	if len(products) == 0 {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "No products found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, products, nil, "")
}
