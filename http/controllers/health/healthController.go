package controllers

import (
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	status := map[string]string{"status": "on"}
	c.JSON(200, status)
}
