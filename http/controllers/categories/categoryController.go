package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to create categories.")
		return
	}

	result := database.PostgresInstance.Create(&newCategory)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new category.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newCategory, nil, "")
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	result := database.PostgresInstance.First(&category, id)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Category not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, category, nil, "")
}

func GetMarketCategories(c *gin.Context) {
	marketSlug := c.Param("slug")

	var market models.Market
	if err := database.PostgresInstance.Where("slug = ?", marketSlug).First(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found.")
		return
	}

	var categories []models.Category
	result := database.PostgresInstance.Where("market_id = ?", market.ID).Find(&categories)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve categories.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, categories, nil, "")
}

func UpdateCategory(c *gin.Context) {
	categoryId := c.Param("id")

	var request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to update categories.")
		return
	}

	var category models.Category

	if err := database.PostgresInstance.First(&category, categoryId).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Category not found.")
		return
	}

	result := database.PostgresInstance.Model(&category).Updates(map[string]interface{}{
		"name":        request.Name,
		"description": request.Description,
	})

	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the category.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, result, nil, "")
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.PostgresInstance.First(&category, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Category not found.")
		return
	}

	if err := database.PostgresInstance.Delete(&category).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete the category.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, nil, nil, "")
}
