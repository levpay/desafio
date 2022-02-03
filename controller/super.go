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
	"strings"

	"github.com/google/uuid"
)

const (
	apiPath = "https://superheroapi.com/api/%s/search/"
)

var (
	apiSearchPath = fmt.Sprintf(apiPath, config.Env.APIKey)
)

func insertUUIDtoSuper(supers models.HeroAPIResponse) []models.SuperInsert {
	var superSlice []models.SuperInsert
	for _, super := range supers.Results {
		newUUID := uuid.New()
		newSuper := models.SuperInsert{
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

func ListHeroes(rw http.ResponseWriter, r *http.Request) {
	args := []string{"alignment"}
	filter := []interface{}{"good"}
	superSlice := data.GetSupers(args, filter)
	superJson := models.CreateListSupersResponse(superSlice)

	rw.Write(superJson)
}

func ListVillains(rw http.ResponseWriter, r *http.Request) {
	args := []string{"alignment"}
	filter := []interface{}{"bad"}
	superSlice := data.GetSupers(args, filter)
	superJson := models.CreateListSupersResponse(superSlice)

	rw.Write(superJson)
}

func SearchForSuper(rw http.ResponseWriter, r *http.Request) {
	var args []string
	var filters []interface{}
	name := r.FormValue("name")
	uuid := r.FormValue("uuid")

	if name != "" {
		args = append(args, "hero_name")
		filters = append(filters, name)
	}
	if uuid != "" {
		args = append(args, "uuid")
		filters = append(filters, uuid)
	}

	superSlice := data.GetSupers(args, filters)
	superJson := models.CreateListSupersResponse(superSlice)

	rw.Write(superJson)
}

func DeleteSuper(rw http.ResponseWriter, r *http.Request) {
	uuids := strings.Split(r.FormValue("uuid"), ",")

	log.Printf("Deleting supers with UUID: %s", strings.Join(uuids, ", "))

	deletedCount, err := data.DeleteSupers(uuids)
	if err != nil {
		log.Printf("Error deleting supers: %s", err)
	}

	if deletedCount == 0 {
		response := models.FailedResponse("No super found")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(response)
		return
	}

	response := models.SuccessfullyDeleted(deletedCount)
	rw.Write(response)
}
