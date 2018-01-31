package main

import (
	"io"
	"net/http"
)

/* HANDS ON # 1 -----------------------------------
1. ListenAndServe on port ":8080" using the default ServeMux.
	- use HandleFunc to add the following routes to the default ServeMux:
		"/"
		"/dog/"
		"/me/
	- add a func for each of the routes.
	- have the "/me/" route print out your name.
----------------------------------- */

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", nilson)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to exercise #1\n - Visit localhost:8080/dog or localhost:8080/me to see more.")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy ruff ruff")
}

func nilson(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Nilson Molina")
}
