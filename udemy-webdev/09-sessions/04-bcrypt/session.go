package main

import (
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// check for session cookie
	c, err := req.Cookie("session")
	if err != nil {
		return user{} // returns empty user
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	// check if user is logged in
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
