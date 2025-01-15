package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg"
)

type CreateUserInput struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	PasswordHash string `json:"password_hash" binding:"required"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.PostgresInstance.Find(&users)
	response.SendGinResponse(c, http.StatusOK, users, nil, "")
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.PostgresInstance.First(&user, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found")
		return
	}
	response.SendGinResponse(c, http.StatusOK, user, nil, "")
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to create user.")
		return
	}

	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.PasswordHash,
		PhoneNumber:  input.PhoneNumber,
		Address:      input.Address,
		City:         input.City,
		State:        input.State,
		ZipCode:      input.ZipCode,
	}

	if err := database.PostgresInstance.Create(&user).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new user.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, user, nil, "")
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.PostgresInstance.First(&user, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found")
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to update user.")
		return
	}
	now := time.Now()
	user.UpdatedAt = &now
	if err := database.PostgresInstance.Save(&user).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to update user.")
		return
	}
	response.SendGinResponse(c, http.StatusOK, user, nil, "")
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.PostgresInstance.First(&user, id).Error; err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found")
		return
	}
	now := time.Now()
	if err := database.PostgresInstance.Model(&user).UpdateColumn("deleted_at", now).Error; err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to delete user.")
		return
	}
	response.SendGinResponse(c, http.StatusOK, gin.H{"message": "User deleted"}, nil, "")
}
