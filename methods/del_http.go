package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	// Creating a new request
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// Modifying the request headers
	req.Header.Set("X-API-KEY", apiKey)

	// Making a new request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < 299 {
		return nil
	}
	return fmt.Errorf("Request to delete user unsuccessful. \n Error %v!", res.StatusCode)
}
