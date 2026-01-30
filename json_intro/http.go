package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func getIssues(url string) ([]Issue, error){
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue

	// Decoding the json into a slice of structs
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&issues)
	if err != nil {
		return nil, err
	}
	return issues, err
}
