package main

import (
	"fmt"
	"net/http"
)

/* ---------------------------
	go run main.go

	in your browser, go to:
	http://localhost:8080/
------------------------------ */

type hotdog int

/* implement the handler interface
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want here")     // to console
	fmt.Fprintln(w, "Any code you want here") // to browser
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
