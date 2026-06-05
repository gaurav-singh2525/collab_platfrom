package middleware

import (
	"net/http"
	"strings"

	"collab-code-platform/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(
	jwtSecret string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader(
			"Authorization",
		)

		if authHeader == "" {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing token",
				},
			)

			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := auth.ValidateToken(
			tokenString,
			jwtSecret,
		)

		if err != nil || !token.Valid {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set(
			"user_id",
			claims["user_id"],
		)

		c.Set(
			"email",
			claims["email"],
		)

		c.Next()
	}
}
