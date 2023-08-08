package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/Tony-Lucas/gopportunities/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserClaim struct {
	jwt.StandardClaims
	ID int
}

type AdminClaim struct {
	jwt.StandardClaims
	ID int
}

func Auth(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"success": "false",
			"message": "Usuário não cadastrado",
		})
	} else {
		result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
		if result != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": "false",
				"message": "Senha inválida",
			})
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
				StandardClaims: jwt.StandardClaims{},
				ID:             int(user.ID),
			})

			jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
			if err != nil {
				log.Fatal(err)
			}
			c.SetCookie("token", jwtToken, 3600, "/", "localhost", false, true)
			c.JSON(200, gin.H{
				"success": "true",
			})
		}
	}
}

func VerifyAuth(c *gin.Context) {
	cookie, err := c.Cookie("token")
	var userClaim UserClaim

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "Token não encontrado",
		})
	} else {
		token, err := jwt.ParseWithClaims(cookie, &userClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"message": err,
			})
		} else {
			if token.Valid {
				c.JSON(200, gin.H{
					"success": "true",
				})
			} else {
				c.JSON(200, gin.H{
					"success": "false",
				})
			}
		}
	}

}
