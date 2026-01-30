package main

import (
	"encoding/json"
)


func marshalAll[T any](items []T) ([][]byte, error) {
	var result [][]byte
	
	/** Marshal converts a go struct into a slice of bytes repping JSON data **/
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	
	return result, nil
}
