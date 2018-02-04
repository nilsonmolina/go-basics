package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/multiple", multiple)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value.",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, " - in chrome, go to:\n\tdev tools / application / cookies")
	fmt.Fprintln(w, " - go to:\n\tlocalhost:8080/multiple\n\tlocalhost:8080/read")
}

func read(w http.ResponseWriter, req *http.Request) {
	for i, c := range req.Cookies() {
		fmt.Fprintf(w, "COOKIE #%v:\n\t%v\n\n", i+1, c)
	}

	// // Informational purposes: error message for a non-existent cookie
	// _, err := req.Cookie("does-not-exist")
	// if err != nil {
	// 	log.Println(err)
	// }
}

func multiple(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things.",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things.",
	})
	fmt.Fprintln(w, "MULTIPLE COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, " - in chrome, go to:\n\tdev tools / application / cookies")
	fmt.Fprintln(w, " - go to:\n\tlocalhost:8080/read")
}
