package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health for checking service status
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
