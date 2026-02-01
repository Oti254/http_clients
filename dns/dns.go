package main

import (
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	// Instantiating a URL struct using url.Parse
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil
}
