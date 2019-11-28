package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet handle get request to ping route
func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"hello": "Found me",
	})
}
