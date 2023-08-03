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
	Username  string `gorm:"not null;unique" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `form:"createdAt" json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `form:"updatedAt" json:"updatedAt" gorm:"not null"`
}

type Product struct {
	gorm.Model
	Name           string `json:"name" binding:"required" gorm:"not null"`
	PriceWholesale string `json:"priceWholesale" binding:"required" gorm:"not null"`
	PriceRetail    string `json:"priceRetail" binding:"required" gorm:"not null"`
}

type DeliverAdress struct {
	gorm.Model
	UserID        uint
	ZipCode       uint   `json:"zipCode" binding:"required" gorm:"not null"`
	StreetAddress string `json:"streetAdress" binding:"required" gorm:"not null"`
	Complement    string `json:"complement" binding:"required"`
	Area          string `json:"area" binding:"required" gorm:"not null"`
	City          string `json:"city" binding:"required" gorm:"not null"`
	State         string `json:"state" binding:"required" gorm:"not null"`
	HouseNumber   uint16 `json:"houseNumber" binding:"required" gorm:"not null"`
}

type Contact struct {
	gorm.Model
	PhoneNumber string `json:"phoneNumber" binding:"required,phoneNumber" gorm:"not null"`
	UserId      uint   `json:"userId" binding:"required,userId" gorm:"not null"`
}
