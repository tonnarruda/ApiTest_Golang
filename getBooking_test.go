package main

import (
	"net/http"
	"testing"
)

func getAllBooking(api *API) string {
	return api.Endpoints["booking"]
}

func TestGetAllBooking(t *testing.T) {
	api := setupApi("")
	endpoint := getAllBooking(api)

	api.Client.Get(endpoint).
		Expect(t).
		Status(http.StatusOK).
		Done()
}
