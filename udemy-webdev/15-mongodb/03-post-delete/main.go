package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nilsonmolina/go-basics/udemy-webdev/15-mongodb/02-models/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en>
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/1323">GO TO: http://localhost/user/1323</a>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK) // STATUS 200
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "Male",
		Age:    43,
		ID:     p.ByName("id"),
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // STATUS 200
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// composite literal - type and curly braces
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from stream
	json.NewDecoder(req.Body).Decode(&u)

	// change ID
	u.ID = "007"

	// marshal/unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // STATUS 201
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // STATUS 200
	fmt.Fprint(w, "Write code to delete user\n")
}
