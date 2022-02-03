package models

import (
	"encoding/json"
	"log"
)

type ListAllSupersResponse struct {
	Status  string  `json:"status"`
	Results []Super `json:"results"`
}

func CreateListSupersResponse(supers []Super) []byte {
	response := ListAllSupersResponse{
		Status:  "success",
		Results: supers,
	}
	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error while creating Supers list response: %s", err)
	}
	return responseByte
}
