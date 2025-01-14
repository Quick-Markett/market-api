package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error string      `json:"error,omitempty"`
}

func SendGinResponse(c *gin.Context, statusCode int, data interface{}, meta interface{}, errMessage string) {
	response := APIResponse{
		Data:  data,
		Meta:  meta,
		Error: errMessage,
	}

	c.JSON(statusCode, response)
}
