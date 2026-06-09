package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		expected := "Bearer " + os.Getenv("ADMIN_TOKEN")
		if token != expected {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
