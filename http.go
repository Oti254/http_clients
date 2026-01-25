package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	// http.Get uses the http.DefaultClient to make a request to the url
	// res is the response from the server
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	// ensures response body is properly closed after reading
	defer res.Body.Close()
	// io.ReadAll reads the response body into a slice of bytes
	data, err  := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return data, err
}

