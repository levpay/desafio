package routes

import (
	"desafio/controller"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/create", controller.CreateSuper)
	router.HandleFunc("/api/listAll", controller.ListAllSupers)
	router.HandleFunc("/api/listHeroes", controller.ListHeroes)
	router.HandleFunc("/api/listVillains", controller.ListVillains)
	router.HandleFunc("/api/search", controller.SearchForSuper)
	router.HandleFunc("/api/delete", controller.DeleteSuper)

	return router
}
