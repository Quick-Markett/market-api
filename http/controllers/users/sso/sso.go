package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func generateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func compareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUserWithSSO(c *gin.Context) {
	var request struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to create user with token.")
		return
	}

	tokenString := request.Token
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("JWT_SECRET")
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		response.SendGinResponse(c, http.StatusUnauthorized, nil, nil, "Invalid or expired token.")
		return
	}

	newUser := models.User{
		Email:          claims["email"].(string),
		Name:           claims["name"].(string),
		ProfilePicture: claims["profile_picture"].(string),
		Password:       claims["password"].(string),
		Address:        claims["address"].(string),
		State:          claims["state"].(string),
		City:           claims["city"].(string),
	}

	hashedPassword, err := generateHash(newUser.Password)
	if err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to hash the password.")
		return
	}

	newUser.Password = hashedPassword

	result := database.PostgresInstance.Create(&newUser)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to create a new user with the provided token.")
		return
	}

	response.SendGinResponse(c, http.StatusCreated, newUser, nil, "")
}

func LoginUserWithSso(c *gin.Context) {
	var request struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid JSON to create user with token.")
		return
	}

	tokenString := request.Token

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("JWT_SECRET")
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		response.SendGinResponse(c, http.StatusUnauthorized, nil, nil, "Invalid or expired token.")
		return
	}

	email, emailOk := claims["email"].(string)
	if !emailOk {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid token claims. Email not found.")
		return
	}

	var user models.User

	result := database.PostgresInstance.Where("email = ?", email).First(&user)
	if result.Error != nil {
		response.SendGinResponse(c, http.StatusNotFound, nil, nil, "User not found.")
		return
	}

	password, passwordOk := claims["password"].(string)
	if !passwordOk {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid token claims. Password not found.")
		return
	}

	if !compareHash(password, user.Password) {
		response.SendGinResponse(c, http.StatusUnauthorized, nil, nil, "Invalid password.")
		return
	}

	newToken, err := generateJWT(user)
	if err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to generate token.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, gin.H{
		"user":  user,
		"token": newToken,
	}, nil, "")
}
