package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1>Intro to Cookies</h1><a href="/set">set a cookie</a>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
	})
	fmt.Fprintln(w, `<h1>Cookie Has Been Set!</h1><a href="/read">read cookie</a>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther) // Status 303
		return
	}
	fmt.Fprintf(w, `<h1>Your Cookie:</h1>%v<br><br><a href="/expire">expire cookie (delete)</a>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther) // Status 303
		return
	}

	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther) //Status 303
}
