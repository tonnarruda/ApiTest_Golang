package main

import (
	"strings"
	"testing"
)

func getBookingById(api *API, uuid string) string {
	return strings.Replace(api.Endpoints["bookingById"], ":id", uuid, -1)
}

func TestGetAllBookingById_InvalidID(t *testing.T) {
	api := setupApi("")
	fakeUUID := "99999999"
	endpoint := getBookingById(api, fakeUUID)

	api.Client.Get(endpoint).
		Expect(t).
		Status(404).
		Done()
}

func TestGetAllBookingById(t *testing.T) {
	api := setupApi("")
	endpoint := getBookingById(api, "1324")

	api.Client.Get(endpoint).
		Expect(t).
		Status(404).
		Done()
}
