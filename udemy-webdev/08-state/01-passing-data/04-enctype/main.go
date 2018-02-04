package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", defaultEnctype)
	http.HandleFunc("/multipart", multipart)
	http.HandleFunc("/text", text)
	http.ListenAndServe(":8080", nil)
}

type data struct {
	Enctype, Description, Body string
}

func defaultEnctype(w http.ResponseWriter, req *http.Request) {
	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	d := data{
		Enctype:     "application/x-www-form-urlencoded",
		Description: "This is the default enctype so you don't typically write it out",
		Body:        body,
	}

	err := tpl.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func multipart(w http.ResponseWriter, req *http.Request) {
	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	d := data{
		Enctype:     "multipart/form-data",
		Description: "In order to upload files, you will need to use this enctype on the html input",
		Body:        body,
	}

	err := tpl.Execute(w, d)
	if err != nil {
		log.Fatalln(err)
	}
}

func text(w http.ResponseWriter, req *http.Request) {
	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	d := data{
		Enctype:     "text/plain",
		Description: "Generally not recommended to be used, stick to the default or multipart enctypes",
		Body:        body,
	}

	err := tpl.Execute(w, d)
	if err != nil {
		log.Fatalln(err)
	}
}
