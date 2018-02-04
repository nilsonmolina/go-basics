package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/* HANDS ON # 1 -----------------------------------
1. using cookies, track how many times a user has been to your website domain.
----------------------------------- */

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", count)
	http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Welcome!\n  - You've visited this site %v time(s)\n\n", cookie.Value)
	fmt.Fprintf(w, "COOKIE: %v\n", cookie)
}

// // ORIGINAL ATTEMPT
// func count(w http.ResponseWriter, req *http.Request) {
// 	c, err := req.Cookie("counter")
// 	if err != nil {
// 		http.SetCookie(w, &http.Cookie{
// 			Name:  "counter",
// 			Value: "1",
// 		})
// 		fmt.Fprintln(w, "Welcome!\n  - This is your first time here!")
// 	} else {
// 		count, _ := strconv.Atoi(c.Value)
// 		count++
// 		http.SetCookie(w, &http.Cookie{
// 			Name:  c.Name,
// 			Value: strconv.Itoa(count),
// 		})
// 		fmt.Fprintf(w, "Welcome Back!\n  - You've visited this site %v times\n", count)
// 	}
// }
