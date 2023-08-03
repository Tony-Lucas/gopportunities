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

func AdminAuth(c *gin.Context){
	var admin models.Admin

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
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
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
}