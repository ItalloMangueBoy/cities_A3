package main

import (
	"acc_api_a3/api"
	"acc_api_a3/config"
	"acc_api_a3/schemas"
	"acc_api_a3/views"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Token":  token.StrToken,
			"Cities": cities,
		}

		if err := views.ListCities().Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
