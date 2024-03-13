package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	CHARGILIY_SECRET_KEY = "test_sk_EOuwxMvBadUUtwFe2WHo0A5asuDtKQPgjkoA56mX"
)

func CreateCheckout(amount int, currency string, webhook_link string) []byte {
	requestBody := map[string]interface{}{
		"amount":           amount,
		"currency":         currency,
		"success_url":      "https://google.com",
		"webhook_endpoint": webhook_link,
	}
	fmt.Println(webhook_link)
	// Marshal the map to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request with the POST method, request body, and URL
	req, err := http.NewRequest(
		"POST",
		"https://pay.chargily.net/test/api/v2/checkouts",
		bytes.NewBuffer(requestBodyJSON),
	)
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+CHARGILIY_SECRET_KEY)

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return response
}
