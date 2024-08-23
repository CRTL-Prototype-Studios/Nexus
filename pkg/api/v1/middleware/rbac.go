package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nexus/pkg/api/v1/models"
	"nexus/pkg/database"
)

func RBACMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userUUID, exists := c.Get("userUUID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		var user models.User
		if err := database.Inst.Preload("Role.Permissions").Where("uuid = ?", userUUID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		hasPermission := false
		for _, permission := range user.Role.Permissions {
			if permission.Name == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
