package main

import (
	"github.com/Tony-Lucas/gopportunities/configuration"
	_ "github.com/Tony-Lucas/gopportunities/models"
	"github.com/Tony-Lucas/gopportunities/router"
	"github.com/joho/godotenv"
)

func main() {

	go configuration.DbSingleInstance()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	router.Initialize()
}
