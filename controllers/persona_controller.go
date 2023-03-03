package controllers

import (
	"apiGoMysql/commons"
	"apiGoMysql/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)
	jsonData, _ := json.Marshal(personas)
	commons.SendResponse(w, http.StatusOK, jsonData)
}

func GetOne(w http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		jsonData, _ := json.Marshal(persona)
		commons.SendResponse(w, http.StatusOK, jsonData)
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}

func SaveOne(w http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	jsonData, _ := json.Marshal(persona)

	commons.SendResponse(w, http.StatusCreated, jsonData)
}

func DeleteOne(w http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(w, http.StatusOK, []byte(`{el registro fue eliminado}`))
	} else {
		commons.SendError(w, http.StatusNotFound)
	}

}
