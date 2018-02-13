package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	s := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>JSON in Go</title>
	</head>
	<body>
		<h1>JSON in Go</h1>
		<a href="/marshal">Marshal</a><br>
		<a href="/encode">Encode</a>
	</body>
	</html>
	`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor", "Marshal"},
	}
	// Marshal takes a pointer to the object and returns a value you can manipulate
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Contenty-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor", "Encode"},
	}
	// Encode does not return a value, it simply sends it over the wire immediately
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
