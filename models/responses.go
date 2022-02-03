package models

import (
	"encoding/json"
	"log"
)

type SuperResponse struct {
	UUID              string     `json:"uuid"`
	Name              string     `json:"name"`
	Powerstats        Powerstats `json:"powerstats"`
	Biography         Biography  `json:"biography"`
	Occupation        string     `json:"occupation"`
	RelativesCount    int        `json:"relatives-count"`
	GroupAffiliations []string   `json:"group-affiliations"`
	Image             string     `json:"image"`
}

type ListAllSupersResponse struct {
	Status  string          `json:"status"`
	Results []SuperResponse `json:"results"`
}

func CreateListSupersResponse(supers []SuperResponse) []byte {
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

func FailedResponse(message string) []byte {
	response := map[string]string{
		"status":  "failed",
		"message": message,
	}
	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error while creating failed response: %s", err)
	}
	return responseByte
}

func SuccessfullyDeleted(count int) []byte {
	response := struct {
		Status       string `json:"status"`
		DeletedCount int    `json:"deleted-count"`
	}{
		Status:       "success",
		DeletedCount: count,
	}

	responseByte, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error while creating success response: %s", err)
	}
	return responseByte
}
