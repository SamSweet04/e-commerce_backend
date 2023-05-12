package middlewares

import (
	"github.com/SamSweet04/e-commerce_backend.git/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthWithJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		if userID, err := auth.ValidateToken(tokenString); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			context.Set("id", userID)
		}
	}
}
