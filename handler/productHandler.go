package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/Tony-Lucas/gopportunities/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
			"product":    product,
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
			"product":    products,
		})
	}

}

func PostProduct(c *gin.Context) {

	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": err,
		})
	}
	if file.Filename == "" {
		product := models.Product{Name: c.PostForm("name"), PriceRetail: c.PostForm("priceRetail"), PriceWholesale: c.PostForm("priceWholesale")}
		result := db.Create(&product)
		if result.Error != nil {
			panic(result.Error)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "true",
				"product":    product,
			})
		}
	} else {

		file.Filename = fmt.Sprint(time.Now().UTC().UnixMilli(), file.Filename)
		f, _ := file.Open()
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_ACCESS_KEY"), ""),
			Region:      aws.String(os.Getenv("BUCKET_REGION"))},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"message": "Ocorreu um erro no sistema",
			})
		}
		uploader := s3manager.NewUploader(sess)
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(os.Getenv("BUCKET_NAME")),
			Key:    aws.String(file.Filename),
			Body:   f,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"message": "Ocorreu um erro no sistema",
			})
		} else {
			imgUrl := "https://go-store-s3.s3.amazonaws.com/" + file.Filename
			product := models.Product{Name: c.PostForm("name"), PriceRetail: c.PostForm("priceRetail"), PriceWholesale: c.PostForm("priceWholesale"), ImgName: file.Filename, ImgUrl: imgUrl}
			result := db.Create(&product)
			if result.Error != nil {
				panic(result.Error)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success": "true",
					"product":    product,
				})
			}
		}

	}

}

func UpdateProduct(c *gin.Context) {

	// get file and put on file variable
	file, _ := c.FormFile("img")

	// get body request and put on product variable
	var product models.Product
	c.Bind(&product)

	if file == nil {

		// update product
		result := db.Save(&product)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"message": result.Error,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "true",
				"product": product,
			})
		}
	} else {

		// change filename putting time miliseconds before original name
		file.Filename = fmt.Sprint(time.Now().UTC().UnixMilli(), file.Filename)

		// Open aws session
		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_ACCESS_KEY"), ""),
			Region:      aws.String(os.Getenv("BUCKET_REGION"))},
		)

		if product.ImgName == "" {

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": "false",
					"message": "Occoreu um erro no sistema",
				})
			} else {
				f, _ := file.Open()

				// create a new uploader to s3
				uploader := s3manager.NewUploader(sess)

				// upload the file to s3
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(os.Getenv("BUCKET_NAME")),
					Key:    aws.String(file.Filename),
					Body:   f,
				})

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": "false",
						"message": "Occoreu um erro no sistema",
					})
				} else {
					product.ImgName = file.Filename
					product.ImgUrl = "https://go-store-s3.s3.amazonaws.com/" + file.Filename
					result := db.Save(&product)
					if result.Error != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"success": "false",
							"message": "Occoreu um erro no sistema",
						})
					} else {
						c.JSON(http.StatusInternalServerError, gin.H{
							"success": "true",
							"product": product,
						})
					}
				}
			}
		} else {

			// Create S3 service client
			svc := s3.New(sess)

			// Request to delete a s3 object
			_, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(os.Getenv("BUCKET_NAME")), Key: aws.String(file.Filename)})

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": "false",
					"message": "Occoreu um erro no sistema",
				})
			} else {
				f, _ := file.Open()

				// create a new uploader to s3
				uploader := s3manager.NewUploader(sess)

				// upload the file to s3
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: aws.String(os.Getenv("BUCKET_NAME")),
					Key:    aws.String(file.Filename),
					Body:   f,
				})
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": "false",
						"message": "Occoreu um erro no sistema",
					})
				} else {
					product.ImgName = file.Filename
					product.ImgUrl = "https://go-store-s3.s3.amazonaws.com/" + file.Filename

					// update product on db
					result := db.Save(&product)

					if result.Error != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"success": "false",
							"message": "Occoreu um erro no sistema",
						})
					} else {
						c.JSON(http.StatusInternalServerError, gin.H{
							"success": "true",
							"product": product,
						})
					}
				}
			}
		}
	}
}

func DeleteProduct(c *gin.Context) {

	var product models.Product
	result := db.First(&product, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": "false",
			"message": "Occoreu um erro no sistema",
		})
	} else {
		if product.ImgName == "" {
			result := db.Delete(&product)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": "false",
					"message": "Occoreu um erro no sistema",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": "true",
				})
			}
		} else {

			// Open aws session
			sess, _ := session.NewSession(&aws.Config{
				Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_ACCESS_KEY"), ""),
				Region:      aws.String(os.Getenv("BUCKET_REGION"))},
			)
			// Create S3 service client
			svc := s3.New(sess)

			// Request to delete a s3 object
			_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(os.Getenv("BUCKET_NAME")), Key: aws.String(product.ImgName)})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": "false",
					"message": "Occoreu um erro no sistema",
				})
			} else {
				db.Delete(&product)
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": "true",
				})
			}
		}
	}

}
