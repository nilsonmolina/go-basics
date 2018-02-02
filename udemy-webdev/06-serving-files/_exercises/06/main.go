package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	executeTpl(w, "index.gohtml")
}

func about(w http.ResponseWriter, req *http.Request) {
	executeTpl(w, "about.gohtml")
}

func contact(w http.ResponseWriter, req *http.Request) {
	executeTpl(w, "contact.gohtml")
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		executeTpl(w, "applyProcess.gohtml")
		return
	}
	executeTpl(w, "apply.gohtml")
}

func executeTpl(w http.ResponseWriter, tplName string) {
	err := tpl.ExecuteTemplate(w, tplName, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
