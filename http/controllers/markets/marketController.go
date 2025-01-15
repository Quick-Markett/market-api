package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/models"
	response "main.go/pkg"
)

type CreateMarketInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Description string `json:"description"`
	LogoUrl     string `json:"logo_url"`
}

func GetMarkets(c *gin.Context) {
	var markets []models.Market
	models.DB.Find(&markets)
	response.SendGinResponse(c, http.StatusOK, markets, nil, "")
}

func GetMarket(c *gin.Context) {
	id := c.Param("id")
	var market models.Market
	if err := models.DB.First(&market, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	response.SendGinResponse(c, http.StatusOK, market, nil, "")
}

func CreateMarket(c *gin.Context) {
	var input CreateMarketInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to create market.")
		return
	}

	market := models.Market{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		City:        input.City,
		State:       input.State,
		ZipCode:     input.ZipCode,
		Description: input.Description,
		LogoUrl:     input.LogoUrl,
	}

	if err := models.DB.Create(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new market.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, market, nil, "")
}

func UpdateMarket(c *gin.Context) {
	id := c.Param("id")
	var market models.Market
	if err := models.DB.First(&market, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	if err := c.ShouldBindJSON(&market); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to update market.")
		return
	}
	if err := models.DB.Save(&market).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update market.")
		return
	}
	response.SendGinResponse(c, http.StatusOK, market, nil, "")
}

func DeleteMarket(c *gin.Context) {
	id := c.Param("id")
	var market models.Market
	if err := models.DB.First(&market, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Market not found")
		return
	}
	now := time.Now()
	if err := models.DB.Model(&market).UpdateColumn("deleted_at", now).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete market.")
		return
	}
	response.SendGinResponse(c, http.StatusOK, gin.H{"message": "Market deleted"}, nil, "")
}
