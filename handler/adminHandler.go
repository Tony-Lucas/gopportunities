package handler

import (
	"fmt"
	"net/http"

	"github.com/Tony-Lucas/gopportunities/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(c *gin.Context) {

	// Hash request password
	hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 10)
	if err != nil {
		panic(err)
	}

	// Initializing a variable with user request data
	admin := &models.Admin{Name: c.PostForm("name"), Username: c.PostForm("username"), Password: string(hash)}

	// Create User db

	if err := db.Create(&admin).Error; err != nil {
		fmt.Println()
	}

	// Initializing a variable with contact request data
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"admin":   admin,
	})
}
