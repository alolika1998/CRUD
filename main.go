package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)
type Person struct {
	Id int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Address *Address `json:"address"`  
}
type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}
var people []Person
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := ranges people {
		if item.Id == params["Id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndpoint(w http.ResopnseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&Person)
	person.Id = params ["Id"] 
	people = append(people, Person)
	json.NewEncoder(w).Encode(people)
}
fund DeletePersonEndpoint(w http.ResponceWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.Id = params["Id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	} 
	json.NewDecoderencoder(w).Encode(people)
}
func main() {
	router := mux.NewRouter()
	people = append(people, Person{Id: 1, Firstname: "Lisa", Lastname: "Roy", Address: &Address{City: "Kolkata", State: "West Begal"}})
	people = append(people, Person{Id: 2, Firstname: "Sam", Lastname: "West"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people{Id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people{Id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people{Id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
