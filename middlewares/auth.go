package middlewares

import (
	"net/http"
	"strings"

	"github.com/faisallbhr/gin-boilerplate/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.ResponseError(c, "Unauthorized", http.StatusUnauthorized, map[string]string{
				"token": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenString == "" {
			helpers.ResponseError(c, "Unauthorized", http.StatusUnauthorized, map[string]string{
				"token": "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		claims, err := helpers.VerifyToken(tokenString)
		if err != nil {
			helpers.ResponseError(c, "Unauthorized", http.StatusUnauthorized, map[string]string{
				"token": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
