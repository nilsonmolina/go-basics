package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "go to:\n - localhost:8080/barred")
	fmt.Printf("Your request method at foo: %v\n\n", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// using http.Redirect
	http.Redirect(w, req, "/", http.StatusSeeOther) // Status 303 - Method becomes GET

	// // using WriteHeader
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
