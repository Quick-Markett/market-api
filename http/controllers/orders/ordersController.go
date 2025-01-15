package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg"
)

func CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to create orders.")
		return
	}

	result := database.PostgresInstance.Create(newOrder)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new order.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newOrder, nil, "")
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	result := database.PostgresInstance.First(&order, id)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Order not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, order, nil, "Order retrieved successfully.")
}

func GetMarketOrders(c *gin.Context) {
	marketID := c.Param("id")
	var orders []models.Order

	result := database.PostgresInstance.Where("market_id = ?", marketID).Find(&orders)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve orders.")
		return
	}

	if len(orders) == 0 {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "No orders found for this market.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, orders, nil, "Market orders retrieved successfully.")
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := database.PostgresInstance.First(&order, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Order not found.")
		return
	}

	var updatedData models.Order
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON for updating the order.")
		return
	}

	result := database.PostgresInstance.Model(&order).Updates(updatedData)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the order.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, order, nil, "Order updated successfully.")
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := database.PostgresInstance.First(&order, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Order not found.")
		return
	}

	if err := database.PostgresInstance.Delete(&order).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete the order.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, nil, nil, "Order deleted successfully.")
}
