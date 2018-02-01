package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<!--image doesn't serve-->
		<h1>Not Serving</h1>
		<h4>Img sourced from this directory w/out serving</h2>
		<img src="../toby.jpg">
		`)

	io.WriteString(w, `
		<!--not serving from our server-->
		<h4> Img sourced from wikipedia, NOT our server</h2>
		<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
		`)
}
