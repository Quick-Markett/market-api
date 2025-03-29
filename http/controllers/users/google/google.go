package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func CreateUserWithGoogle(c *gin.Context) {
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

func LoginUserWithGoogle(c *gin.Context) {
	googleId := c.Param("googleId")
	var user models.User

	result := database.PostgresInstance.Where("google_id = ?", googleId).First(&user)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, user, nil, "")
}
