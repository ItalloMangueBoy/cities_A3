package main

import (
	"acc_api_a3/api"
	"acc_api_a3/config"
	"acc_api_a3/schemas"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config.InitClient()

	auth := schemas.Auth{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}

	token, err := api.Login(auth)
	if err != nil {
		log.Fatal(err)
	}

	cities, err := api.GetCities(token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(cities)
}
