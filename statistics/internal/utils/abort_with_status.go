package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AbortWithStatusCode(c *gin.Context, statusCode int, msg string, err error) {
	c.AbortWithStatusJSON(statusCode, GetResponse(fmt.Sprintf("%s %s", msg, err)))
}
