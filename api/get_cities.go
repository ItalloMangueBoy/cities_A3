package api

import (
	"acc_api_a3/config"
	"acc_api_a3/schemas"
	"encoding/json"
)

func GetCities(token schemas.Token) ([]schemas.City, error) {
	res, err := config.Client.R().
		SetHeader("x-access-token", token.StrToken).
		Get("/cidades")
	if err != nil {
		return nil, err
	}

	var cities []schemas.City
	err = json.Unmarshal(res.Body(), &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
