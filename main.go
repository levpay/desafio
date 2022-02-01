package main

import (
	"desafio/config"
	"desafio/routes"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	router := routes.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	log.Println("Server starting on port: " + config.Env.ServerPort)
	http.ListenAndServe(":"+config.Env.ServerPort, n)

}
