package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "go to this location twice:\n - localhost:8080/bar")
	fmt.Printf("Your request method at foo: %v\n\n", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently) // Status 301 -
}
