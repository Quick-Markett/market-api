package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "main.go/models"
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
    c.JSON(http.StatusOK, markets)
}

func GetMarket(c *gin.Context) {
    id := c.Param("id")
    var market models.Market
    if err := models.DB.First(&market, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Market not found"})
        return
    }
    c.JSON(http.StatusOK, market)
}

func CreateMarket(c *gin.Context) {
    var input CreateMarketInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

    models.DB.Create(&market)
    c.JSON(http.StatusCreated, market)
}

func UpdateMarket(c *gin.Context) {
    id := c.Param("id")
    var market models.Market
    if err := models.DB.First(&market, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Market not found"})
        return
    }
    if err := c.ShouldBindJSON(&market); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    now := time.Now()
    market.UpdatedAt = &now
    models.DB.Save(&market)
    c.JSON(http.StatusOK, market)
}

func DeleteMarket(c *gin.Context) {
    id := c.Param("id")
    var market models.Market
    if err := models.DB.First(&market, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Market not found"})
        return
    }
    now := time.Now()
    models.DB.Model(&market).UpdateColumn("deleted_at", now)
    c.JSON(http.StatusOK, gin.H{"message": "Market deleted"})
}