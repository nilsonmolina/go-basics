package main

import (
	"html/template"
	"log"
	"net/http"
)

/* HANDS ON # 2 -----------------------------------
1. Take the previous program in the previous folder and change it so that:
	- a template is parsed and served
	- you pass data into the template
----------------------------------- */

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", nilson)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	s := "Welcome to exercise #2 - visit localhost:8080/dog/ or localhost:8080/me/ to see more."
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	handleErr(err)
}

func dog(w http.ResponseWriter, req *http.Request) {
	s := "doggy doggy ruff ruff"
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	handleErr(err)
}

func nilson(w http.ResponseWriter, req *http.Request) {
	s := "Nilson Molina"
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
