package main

import (
	"apiGoMysql/commons"
	"apiGoMysql/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	log.Println("Servidor ejecutandose sobre el puerto :3000")
	log.Println(server.ListenAndServe())
}
