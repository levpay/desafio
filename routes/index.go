package routes

import (
	"desafio/controller"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/create", controller.CreateSuper)
	router.HandleFunc("/api/listAll", controller.ListAllSupers)

	return router
}
