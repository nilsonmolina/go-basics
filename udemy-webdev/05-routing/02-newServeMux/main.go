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

	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // can catch /dog && /dog/*
	mux.Handle("/cat", c)  // can catch /cat but NOT /cat/*

	http.ListenAndServe(":8080", mux)
}
