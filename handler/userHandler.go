package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tony-Lucas/gopportunities/configuration"
	"github.com/Tony-Lucas/gopportunities/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	User    models.User
	Contact models.Contact
}

var db = configuration.DbSingleInstance()

func GetUser(c *gin.Context) {

	var user *models.User
	db.First(&user, c.Param("id"))
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {

	// Hash request password
	hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), 10)
	if err != nil {
		panic(err)
	}

	// Initializing a variable with user request data
	contacts := []models.Contact{{PhoneNumber: c.PostForm("phoneNumber")}}
	user := &models.User{Name: c.PostForm("name"), Lastname: c.PostForm("lastname"), Email: c.PostForm("email"), Password: string(hash), Contacts: contacts}

	// Create User db

	if err := db.Create(&user).Error; err != nil {
		fmt.Println()
	}

	// Initializing a variable with contact request data
	json, _ := json.Marshal(user)
	c.JSON(200, json)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	db.Model(&user).Updates(models.User{Name: user.Name, Email: user.Email, Lastname: user.Lastname})
	c.JSON(http.StatusOK, user)
}
