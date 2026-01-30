package main

import (
	"io"
	"encoding/json"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var issues []Issue

	// json.Decoder streams data from io.ReadAll
	// json.Unmarshal works with data already in the []byte format
	// json.Unmarshal is ideal for small amounts of data
	// json.Decoder is more memory efficient since it doesn't load all data at once into memory
	err = json.Unmarshal(data, &issues)
	if err != nil {
		return nil, err
	}
	return issues, err
}
