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
	var admin models.Admin
	if c.PostForm("username") != "" {
		if result := db.Where("username = ?", c.PostForm("username")).First(&admin); result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"message": "Usuário não cadastrado",
			})
		} else {
			result := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(c.PostForm("password")))
			if result != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"success": "false",
					"message": "Senha inválida",
				})
			} else {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, AdminClaim{
					StandardClaims: jwt.StandardClaims{},
					ID:             int(admin.ID),
				})

				jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY_ADMIN")))
				if err != nil {
					log.Fatal(err)
				}
				c.SetCookie("token", jwtToken, 3600, "/", "localhost", false, true)
				c.JSON(200, gin.H{
					"success": "true",
				})
			}
		}
	} else if c.PostForm("email") != "" {
		if result := db.Where("email = ?", c.PostForm("email")).First(&user); result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": "false",
				"message": "Usuário não cadastrado",
			})
		} else {
			result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
			if result != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
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
				c.SetCookie("token", jwtToken, 3600, "/", "gopportunities-production-8428.up.railway.app", true, true)
				c.JSON(200, gin.H{
					"success": "true",
				})
			}
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
