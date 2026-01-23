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
	convertedBytes := string(issues)
	fmt.Println(convertedBytes)
}
