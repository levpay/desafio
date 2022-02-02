package controller

import (
	"desafio/config"
	"desafio/data"
	"desafio/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const (
	apiPath = "https://superheroapi.com/api/%s/search/"
)

var (
	apiSearchPath = fmt.Sprintf(apiPath, config.Env.APIKey)
)

func insertUUIDtoSuper(supers models.HeroAPIResponse) []models.Super {
	var superSlice []models.Super
	for _, super := range supers.Results {
		newUUID := uuid.New()
		newSuper := models.Super{
			HeroAPIResults: super,
			UUID:           newUUID.String(),
		}
		superSlice = append(superSlice, newSuper)
	}
	return superSlice
}

func CreateSuper(rw http.ResponseWriter, r *http.Request) {
	superName := r.FormValue("name")

	apiResponse, err := http.Get(apiSearchPath + superName)
	if err != nil {
		log.Println("Error accessing external API")
	}

	var heroAPIResponse models.HeroAPIResponse
	responseData, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Println("Error reading API response")
	}
	json.Unmarshal(responseData, &heroAPIResponse)

	supers := insertUUIDtoSuper(heroAPIResponse)

	err = data.InsertSuper(supers)
	if err != nil {
		log.Printf("Error inserting to database: %s\n", err)
	}
	log.Printf("Supers inserted: %v\n", len(supers))

	response := map[string]string{
		"status":        "success",
		"created-count": strconv.Itoa(len(supers)),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Println("Error unmarshaling response")
		rw.Write([]byte("Error! Could not insert to database"))
		return
	}

	rw.Write(responseJSON)

}

func ListAllSupers(rw http.ResponseWriter, r *http.Request) {
	superSlice := data.GetSupers(nil, nil)
	supersJson := models.CreateListSupersResponse(superSlice)

	rw.Write(supersJson)
}
