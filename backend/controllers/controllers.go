package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gilsondev/go-rest-api/database"
	"github.com/gilsondev/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	var personalidades []models.Personalidade

	database.DB.Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	if personalidade.Id > 0 {
		json.NewEncoder(w).Encode(personalidade)
		return
	}

	http.NotFound(w, r)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
  var personalidade models.Personalidade
  json.NewDecoder(r.Body).Decode(&personalidade)

  database.DB.Create(&personalidade)

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(personalidade)
}

func AtualizaPersonalidade(w http.ResponseWriter, r *http.Request) {
  var personalidade models.Personalidade

  vars := mux.Vars(r)
  id := vars["id"]

  database.DB.First(&personalidade, id)

  if personalidade.Id == 0 {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  json.NewDecoder(r.Body).Decode(&personalidade)
  database.DB.Save(&personalidade)
  json.NewEncoder(w).Encode(personalidade)
  
}

func RemovePersonalidade(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  var personalidade models.Personalidade

  err := database.DB.Select("id").Where("id = ?", id).First(&personalidade).Error

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  database.DB.Delete(&personalidade, &id)
}
