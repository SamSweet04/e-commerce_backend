package middlewares

import (
	"github.com/SamSweet04/e-commerce_backend.git/auth"
	"github.com/gin-gonic/gin"
)

func AuthWithJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		if userID, err := auth.ValidateToken(tokenString); err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
		} else {
			context.Set("userID", userID)
		}
	}
}
