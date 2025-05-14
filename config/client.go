package config

import (
	"os"

	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func InitClient() {
	Client = resty.New()
	Client.SetBaseURL(os.Getenv("API_URL"))
	Client.SetHeader("Content-Type", "application/json")
	Client.SetHeader("Accept", "*/*")
	Client.SetHeader("Accept-Encoding", "gzip, deflate")
	Client.SetHeader("Cache-Control", "no-cache")
	Client.SetHeader("Connection", "keep-alive")
}
