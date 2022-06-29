package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruyoutor/go-api-rest/database"
	"github.com/ruyoutor/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {

	var p []models.Personality
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var personality models.Personality
	database.DB.Find(&personality, id)
	json.NewEncoder(w).Encode(personality)

}
