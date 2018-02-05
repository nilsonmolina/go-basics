package main

import (
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// check for session - cookie
	c, err := req.Cookie("session")
	if err != nil {
		return user{}
	}

	// if user exists, get user
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
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
