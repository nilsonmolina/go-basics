package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person :
type Person struct {
	ID        string  `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Address   Address `json:"address"` // `json:"address,omitempty"`
}

// Address :
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []Person

func init() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Person{})
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
