package getParams

import "github.com/gin-gonic/gin"

func GetParams(c *gin.Context, key string) (string, bool) {
	value := c.Query(key)
	return value, value != ""
}