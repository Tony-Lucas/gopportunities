package handler

import (
	"net/http"

	"github.com/Tony-Lucas/gopportunities/models"
	"github.com/gin-gonic/gin"
)

func GetUserAdress(c *gin.Context) {
	var adress *models.DeliverAdress
	result := db.First(&adress, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"adress":  adress,
		})
	}
}

func GetUserAdresses(c *gin.Context) {
	var adresses []*models.DeliverAdress
	result := db.Where("userId = ?", c.Param("userId")).Find(&adresses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success":  "true",
			"adresses": adresses,
		})
	}
}

func PostUserAdress(c *gin.Context) {
	var adress *models.DeliverAdress
	c.Bind(&adress)
	result := db.Create(&adress)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"adress":  adress,
		})
	}
}

func UpdateUserAdress(c *gin.Context) {
	var adress *models.DeliverAdress
	c.Bind(&adress)
	result := db.Save(&adress)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"adress":  adress,
		})
	}
}

func DeleteUserAdress(c *gin.Context) {
	var adress *models.DeliverAdress
	result := db.Delete(&adress, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
		})
	}
}
