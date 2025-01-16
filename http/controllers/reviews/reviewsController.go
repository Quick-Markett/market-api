package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func CreateReview(c *gin.Context) {
	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to create reviews.")
		return
	}

	result := database.PostgresInstance.Create(newReview)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new review.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newReview, nil, "")
}

func GetReviewById(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	result := database.PostgresInstance.First(&review, id)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Review not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, review, nil, "Review retrieved successfully.")
}

func GetOrderReview(c *gin.Context) {
	marketID := c.Param("id")
	var reviews []models.Review

	result := database.PostgresInstance.Where("market_id = ?", marketID).Find(&reviews)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve reviews.")
		return
	}

	if len(reviews) == 0 {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "No reviews found for this market.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, reviews, nil, "Market reviews retrieved successfully.")
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := database.PostgresInstance.First(&review, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Review not found.")
		return
	}

	var updatedData models.Review
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON for updating the review.")
		return
	}

	result := database.PostgresInstance.Model(&review).Updates(updatedData)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the review.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, review, nil, "Review updated successfully.")
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := database.PostgresInstance.First(&review, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Review not found.")
		return
	}

	if err := database.PostgresInstance.Delete(&review).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete the review.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, nil, nil, "Review deleted successfully.")
}
