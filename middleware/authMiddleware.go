package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserClaim struct {
	jwt.StandardClaims
	ID int
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("token")
		if cookie == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": "false",
				"message": "Token inexistente ou inválido",
			})
		}
		var userClaim UserClaim
		token, _ := jwt.ParseWithClaims(cookie, &userClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": "false",
			})
		} else {
			c.Next()
		}
	}
}
