package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func CreateUserWithSSO(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "Invalid JSON to create user with google AuthO2.")
		return
	}

	result := database.PostgresInstance.Create(&newUser)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new user with google AuthO2.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newUser, nil, "")
}

func LoginUserWithSso(c *gin.Context) {
	email := c.Param("email")
	var user models.User

	result := database.PostgresInstance.Where("email = ?", email).First(&user)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, user, nil, "")
}
