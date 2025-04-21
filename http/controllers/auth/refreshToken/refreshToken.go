package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	response "main.go/pkg/response"
)

func RefreshToken(c *gin.Context) {
	var request struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.SendGinResponse(c, http.StatusBadRequest, nil, nil, "Invalid token claims. Email not found.")
		return
	}

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		response.SendGinResponse(c, http.StatusUnauthorized, nil, nil, "Invalid Token.")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to generate token.")
		return
	}

	claims["exp"] = time.Now().Add(120 * time.Minute).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response.SendGinResponse(c, http.StatusInternalServerError, nil, nil, "Failed to generate token.")
		return
	}

	response.SendGinResponse(c, http.StatusOK, gin.H{
		"token": signedToken,
	}, nil, "")
}
