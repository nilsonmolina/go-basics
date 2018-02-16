package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nilsonmolina/go-basics/udemy-webdev/15-mongodb/04-controllers/models"
)

// UserController handles all user calls
type UserController struct{}

// NewUserController returns the address to a UserController
func NewUserController() *UserController {
	return &UserController{}
}

// GetUser returns a user in JSON
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		ID:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // STATUS 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser will create a new user
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)

	u.ID = "007"

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Contenty-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // STATYS 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser deletes a user by given ID
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
