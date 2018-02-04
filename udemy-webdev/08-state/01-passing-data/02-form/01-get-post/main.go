package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// FORM POST: sends data through the body (pay load)
	io.WriteString(w, `
		<form method="post">
			<input type="text" name="q">
			<input type="submit">
		</form>
		<br>`+v)

	// // FORM GET: sends data through url
	// io.WriteString(w, `
	// 	<form method="get">
	// 		<input type="text" name="q">
	// 		<input type="submit">
	// 	</form>
	// 	<br>`+v)
}
