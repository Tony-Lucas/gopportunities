package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Tony-Lucas/gopportunities/models"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var product models.Product
	result := db.First(&product, c.Param("id"))
	if result.Error != nil {
		panic(result.Error)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"user":    product,
		})
	}

}

func GetProducts(c *gin.Context) {
	var products []models.Product
	o, _ := strconv.Atoi(c.Param("offset"))
	result := db.Limit(10).Offset(o).Find(&products)
	if result.Error != nil {
		panic(result.Error)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"user":    products,
		})
	}

}

func PostProduct(c *gin.Context) {
	product := models.Product{Name: c.PostForm("name"), PriceRetail: c.PostForm("priceRetail"), PriceWholesale: c.PostForm("priceWholesale")}
	result := db.Create(&product)
	if result.Error != nil {
		panic(result.Error)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"user":    product,
		})
	}
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	c.Bind(&product)
	fmt.Println(product)
	db.Save(&product)
}

func DeleteProduct(c *gin.Context) {
	var product *models.Product
	result := db.Delete(&product,c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": result.Error,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"success":  "true",
		})
	}
}
