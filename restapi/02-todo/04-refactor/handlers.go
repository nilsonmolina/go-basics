package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

// Index :
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

// TodoIndex :
func TodoIndex(w http.ResponseWriter, req *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow :
func TodoShow(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoID)
}
