package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AdminClaim struct {
	jwt.StandardClaims
	ID int
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("token")
		if cookie == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": "false",
				"message": "Token inexistente ou inválido",
			})
		}
		var AdminClaim AdminClaim
		token, _ := jwt.ParseWithClaims(cookie, &AdminClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY_ADMIN")), nil
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
