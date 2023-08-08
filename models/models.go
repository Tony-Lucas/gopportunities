package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint      `gorm:"primaryKey" json:"id" form:"id"`
	CreatedAt       time.Time `form:"createdAt" json:"createdAt" gorm:"not null"`
	UpdatedAt       time.Time `form:"updatedAt" json:"updatedAt" gorm:"not null"`
	DeletedAt       gorm.DeletedAt
	Name            string          `form:"name" json:"name" binding:"required" gorm:"not null"`
	Lastname        string          `form:"lastname" json:"lastname" binding:"required" gorm:"not null"`
	Email           string          `form:"email" json:"email" binding:"required" gorm:"unique;not null"`
	Password        string          `form:"password" json:"password" binding:"required" gorm:"not null"`
	DeliverAdresses []DeliverAdress `form:"deliverAdresses" json:"deliverAdresses"`
	Contacts        []Contact       `form:"contacts" json:"contacts"`
}

type Admin struct {
	ID        uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name      string `gorm:"not null" json:"name"`
	Username  string `gorm:"not null;unique" json:"username" form:"username"`
	Password  string `gorm:"not null" json:"password" form:"password"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `form:"createdAt" json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `form:"updatedAt" json:"updatedAt" gorm:"not null"`
}

type Product struct {
	ID             uint   `gorm:"primaryKey" json:"id" form:"id"`
	Name           string `json:"name" gorm:"not null" form:"name"`
	PriceWholesale string `json:"priceWholesale" gorm:"not null" form:"priceWholesale"`
	PriceRetail    string `json:"priceRetail" gorm:"not null" form:"priceRetail"`
	ImgName        string `json:"imgName" form:"imgName"`
	ImgUrl         string `json:"imgUrl" form:"imgUrl"`
	DeletedAt      gorm.DeletedAt
	CreatedAt      time.Time `form:"createdAt" json:"createdAt" gorm:"not null"`
	UpdatedAt      time.Time `form:"updatedAt" json:"updatedAt" gorm:"not null"`
}

type DeliverAdress struct {
	ID            uint   `gorm:"primaryKey" json:"id" form:"id"`
	UserID        uint   `json:"userId" gorm:"not null" form:"userId"`
	ZipCode       uint   `json:"zipCode" gorm:"not null" form:"zipCode"`
	StreetAddress string `json:"streetAdress" gorm:"not null" form:"streetAdress"`
	Complement    string `json:"complement" form:"complement" `
	Area          string `json:"area"  gorm:"not null" form:"area"`
	City          string `json:"city"  gorm:"not null" form:"city"`
	State         string `json:"state"  gorm:"not null" form:"state"`
	HouseNumber   uint16 `json:"houseNumber" gorm:"not null" form:"houseNumber"`
	DeletedAt     gorm.DeletedAt
	CreatedAt     time.Time `form:"createdAt" json:"createdAt" gorm:"not null"`
	UpdatedAt     time.Time `form:"updatedAt" json:"updatedAt" gorm:"not null"`
}

type Contact struct {
	ID          uint   `gorm:"primaryKey" json:"id" form:"id"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null" form:"phoneNumber"`
	UserId      uint   `json:"userId" gorm:"not null" form:"userId"`
	DeletedAt   gorm.DeletedAt
	CreatedAt   time.Time `form:"createdAt" json:"createdAt" gorm:"not null"`
	UpdatedAt   time.Time `form:"updatedAt" json:"updatedAt" gorm:"not null"`
}
