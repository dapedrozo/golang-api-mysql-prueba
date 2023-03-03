package routes

import (
	"apiGoMysql/controllers"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveOne).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteOne).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}", controllers.GetOne).Methods("GET")
}
