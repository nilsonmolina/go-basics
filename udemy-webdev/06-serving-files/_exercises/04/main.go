package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	pics := http.StripPrefix("/resources", http.FileServer(http.Dir("./public")))
	http.Handle("/resources/", pics)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(w, nil)
}
