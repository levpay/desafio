package controller

import (
	"desafio/config"
	"desafio/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiPath = "https://superheroapi.com/api/%s/search/"
)

var (
	apiSearchPath = fmt.Sprintf(apiPath, config.Env.APIKey)
)

func CreateSuper(rw http.ResponseWriter, r *http.Request) {
	superName := r.FormValue("name")

	apiResponse, err := http.Get(apiSearchPath + superName)
	if err != nil {
		log.Println("Error accessing external API")
	}

	var response models.HeroAPIResponse
	data, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Println("Error reading API response")
	}

	json.Unmarshal(data, &response)
	log.Printf("Returned objects: %v", len(response.Results))

	stringResponse, _ := json.Marshal(response)

	rw.Write([]byte(stringResponse))
}
