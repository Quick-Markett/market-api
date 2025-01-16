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

	result := database.PostgresInstance.Create(newCategory)
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

	response.SendGinResponse(c, http.StatusOK, category, nil, "Category retrieved successfully.")
}

func GetMarketCategories(c *gin.Context) {
	marketID := c.Param("id")
	var categories []models.Category

	result := database.PostgresInstance.Where("market_id = ?", marketID).Find(&categories)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve categories.")
		return
	}

	if len(categories) == 0 {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "No categories found for this market.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, categories, nil, "Market categories retrieved successfully.")
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.PostgresInstance.First(&category, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Category not found.")
		return
	}

	var updatedData models.Category
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON for updating the category.")
		return
	}

	result := database.PostgresInstance.Model(&category).Updates(updatedData)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the category.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, category, nil, "Category updated successfully.")
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

	response.SendGinResponse(c, http.StatusOK, nil, nil, "Category deleted successfully.")
}
