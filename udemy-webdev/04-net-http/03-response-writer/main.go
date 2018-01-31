package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h2>Check devtools and look at the Response Header.</h2>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
