package middlewares

import (
	"github.com/SamSweet04/e-commerce_backend.git/database"
	"github.com/SamSweet04/e-commerce_backend.git/models"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		userID, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"msg": "User hasn't logged in yet"})
			return
		}
		var user models.User
		database.DB.First(&user, userID)

		// Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to load policy from DB"})
			return
		}

		// Casbin enforces policy
		ok, err := enforcer.Enforce(user.Role, obj, act)

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"msg": "Error occurred when authorizing user"})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"msg": "You are not authorized"})
			return
		}
		c.Next()
	}
}
