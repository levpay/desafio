package models

import (
	"encoding/json"
	"log"
)

type ListAllSupersResponse struct {
	status string
	results []Super
}

func CreateListSupersResponse(supers []Super) []byte {
	response := ListAllSupersResponse{
		status: "success",
		results: supers,
	}
	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error while creating Supers list response: %s", err)
	}
	return responseByte
}