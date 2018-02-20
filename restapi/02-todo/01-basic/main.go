package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type todos []todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

func todoIndex(w http.ResponseWriter, req *http.Request) {
	todos := todos{
		todo{Name: "Write presentation"},
		todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func todoShow(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoID)
}
