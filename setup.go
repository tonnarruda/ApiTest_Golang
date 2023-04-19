package main

import (
	"gopkg.in/h2non/baloo.v3"
)

type API struct {
	Client    *baloo.Client
	Endpoints map[string]string
}

func setupApi(token string) *API {
	baseURL := "https://restful-booker.herokuapp.com"

	endpoints := map[string]string{
		"booking":     "/booking",
		"bookingById": "/booking/:id",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	api := baloo.New(baseURL).
		SetHeaders(headers)

	return &API{
		Client:    api,
		Endpoints: endpoints,
	}
}
