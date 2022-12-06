package routes

import (
	db "crudapp/database"
	. "crudapp/models"
	"encoding/json"

	"net/http"


	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func GetPersonList(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personList := db.GetAllPerson()
	json.NewEncoder(w).Encode(personList)
}

func GetPerson(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	personList := db.GetPerson(params["id"])
	if len(personList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No person found with given id found")
	} else {
		json.NewEncoder(w).Encode(personList[0])
	}
}

func CreatePerson(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	result := db.AddPerson(person)

	if result {
		json.NewEncoder(w).Encode("Created successfully")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not create person")
	}
}

func UpdatePerson(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	result := db.UpdatePerson(params["id"], person)

	if result {
		json.NewEncoder(w).Encode("Updated successfully")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not update person")
	}
}

func DeletePerson(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := db.DeletePerson(params["id"])

	if result {
		json.NewEncoder(w).Encode("Deleted successfully")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could not delete person")
	}
}



func RegisterPersonRoutes(r* mux.Router) {
	r.HandleFunc("/api/person", GetPersonList).Methods("GET")
	r.HandleFunc("/api/person/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/api/person", CreatePerson).Methods("POST")
	r.HandleFunc("/api/person/{id}", UpdatePerson).Methods("PATCH")
	r.HandleFunc("/api/person/{id}", DeletePerson).Methods("DELETE")
	
}