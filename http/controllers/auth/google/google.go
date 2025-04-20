package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	response "main.go/pkg/response"
)

func generateJWT(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

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
	type GoogleLoginPayload struct {
		GoogleId string `json:"google_id"`
	}

	var payload GoogleLoginPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to create user with google AuthO2.")
		return
	}

	var user models.User

	result := database.PostgresInstance.Where("google_id = ?", payload.GoogleId).First(&user)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found.")
		return
	}

	fmt.Println("User ID:", user.ID)

	tokenString, err := generateJWT(user)
	if err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to generate JWT token.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	}, nil, "")
}
