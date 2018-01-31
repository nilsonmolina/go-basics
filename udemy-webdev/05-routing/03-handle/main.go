package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dogggy")
}

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kittttyyyy")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d) // can catch /dog && /dog/*
	http.Handle("/cat", c)  // can catch /cat but NOT /cat/*

	// passing in a handler of nil means the default server mux is being used
	http.ListenAndServe(":8080", nil)
}
