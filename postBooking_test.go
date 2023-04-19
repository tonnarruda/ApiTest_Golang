package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/bxcodec/faker"
)

var jsonBody map[string]interface{}

func init() {
	jsonBody = map[string]interface{}{
		"firstname":   faker.FirstName(),
		"lastname":    faker.LastName(),
		"totalprice":  111,
		"depositpaid": true,
		"bookingdates": map[string]interface{}{
			"checkin":  time.Now().AddDate(0, 0, 5).Format("2006-01-02"),
			"checkout": time.Now().AddDate(0, 0, 12).Format("2006-01-02"),
		},
		"additionalneeds": "Breakfast",
	}
}

func postBooking(api *API) string {
	return api.Endpoints["booking"]
}

func TestPostBooking(t *testing.T) {
	api := setupApi("")
	endpoint := postBooking(api)

	fmt.Println(endpoint)

	api.Client.Post(endpoint).
		JSON(jsonBody).
		Expect(t).
		Status(http.StatusOK).
		Done()
}
