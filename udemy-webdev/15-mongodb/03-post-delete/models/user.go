package models

// User : the json tags allow us to rename when marshaling
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	ID     string `json:"id"`
}
