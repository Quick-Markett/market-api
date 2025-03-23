package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func GetMarkets(c *gin.Context) {
	var markets []models.Market
	database.PostgresInstance.Find(&markets)
	response.SendGinResponse(c, http.StatusOK, markets, nil, "")
}

func GetMarket(c *gin.Context) {
	id := c.Param("id")
	var market models.Market
	if err := database.PostgresInstance.Preload("User").First(&market, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	response.SendGinResponse(c, http.StatusOK, market, nil, "")
}

func GetMarketBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var market models.Market
	if err := database.PostgresInstance.Preload("User").Where("slug = ?", slug).First(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	response.SendGinResponse(c, http.StatusOK, market, nil, "")
}

func CreateMarket(c *gin.Context) {
	var newMarket models.Market

	if err := c.ShouldBindJSON(&newMarket); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to create market.")
		return
	}

	if err := database.PostgresInstance.Create(&newMarket).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new market.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newMarket, nil, "")
}

func UpdateMarket(c *gin.Context) {
	id := c.Param("id")
	var market models.Market
	if err := database.PostgresInstance.First(&market, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	if err := c.ShouldBindJSON(&market); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to update market.")
		return
	}
	if err := database.PostgresInstance.Save(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update market.")
		return
	}
	response.SendGinResponse(c, http.StatusOK, market, nil, "")
}

func DeleteMarket(c *gin.Context) {
	id := c.Param("id")
	var market models.Market
	if err := database.PostgresInstance.First(&market, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	now := time.Now()
	if err := database.PostgresInstance.Model(&market).UpdateColumn("deleted_at", now).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete market.")
		return
	}
	response.SendGinResponse(c, http.StatusOK, gin.H{"message": "Market deleted"}, nil, "")
}
