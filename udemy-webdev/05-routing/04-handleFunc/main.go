package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dogy dogy dogggyyyy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kittttyyyy")
}

func main() {
	http.HandleFunc("/dog/", d) // can catch /dog && /dog/*
	http.HandleFunc("/cat", c)  // can catch /cat but NOT /cat/*

	// Alternative - handlerFunc
	// http.Handle("/dog/", http.HandlerFunc(d))
	// http.Handle("/cat", http.HandlerFunc(c))

	// passing in a handler of nil means the 'default server mux' is being used
	http.ListenAndServe(":8080", nil)
}
