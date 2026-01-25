package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data %v", err)
	}
	// Converts a slice of bytes into a string
	convertBytes := string(issues)
	fmt.Println(convertBytes)
}
