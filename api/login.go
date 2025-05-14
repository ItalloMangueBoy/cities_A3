package api

import (
	"acc_api_a3/config"
	"acc_api_a3/schemas"
	"bytes"
	"encoding/json"
)

func Login(form schemas.Auth) (schemas.Token, error) {
	data := new(bytes.Buffer)

	err := json.NewEncoder(data).Encode(&form)
	if err != nil {
		return schemas.Token{}, err
	}

	res, err := config.Client.R().
		SetBody(data).
		Post("/login")
	if err != nil {
		return schemas.Token{}, err
	}

	var token schemas.Token
	err = json.Unmarshal(res.Body(), &token)
	if err != nil {
		return schemas.Token{}, err
	}

	return token, nil
}
