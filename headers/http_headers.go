package main

import (
	"net/http"
)

// reading the Content-Type header from the response
func getContentType(res *http.Response) (string) {
	return res.Header.Get("Content-Type")
	
}
