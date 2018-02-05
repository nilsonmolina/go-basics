package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err.Error())
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Secure: true, // requires secure connection - https
			HttpOnly: true, // prevents javascript access to cookie
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
