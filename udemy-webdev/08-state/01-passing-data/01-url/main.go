package main

import (
	"io"
	"net/http"
)

// visit this page
// http://localhost:8080/?q=dog

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Do my search:\n - "+v)
}
