package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func CreateOrderItem(c *gin.Context) {
	var newOrderItem models.OrderItem
	if err := c.ShouldBindJSON(&newOrderItem); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to create order items.")
		return
	}

	result := database.PostgresInstance.Create(newOrderItem)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new order item.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newOrderItem, nil, "")
}

func GetOrderItemById(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	result := database.PostgresInstance.First(&orderItem, id)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Order item Item not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, orderItem, nil, "Order item retrieved successfully.")
}

func GetOrderItems(c *gin.Context) {
	order := c.Param("id")
	var items []models.OrderItem

	result := database.PostgresInstance.Where("market_id = ?", order).Find(&items)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to retrieve order items.")
		return
	}

	if len(items) == 0 {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "No order items found for this market.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, items, nil, "Order items retrieved successfully.")
}

func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := database.PostgresInstance.First(&orderItem, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Order item not found.")
		return
	}

	var updatedData models.OrderItem
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON for updating the order item.")
		return
	}

	result := database.PostgresInstance.Model(&orderItem).Updates(updatedData)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update the order item.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, orderItem, nil, "Order item updated successfully.")
}

func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := database.PostgresInstance.First(&orderItem, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Order item not found.")
		return
	}

	if err := database.PostgresInstance.Delete(&orderItem).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete the order item.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, nil, nil, "Order item deleted successfully.")
}
