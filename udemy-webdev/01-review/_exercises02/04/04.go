package main

import (
	"fmt"
)

/* QUESTION 4 -----------------------------
	- using the struct 'person', using a composite literal, create a value
		of type person and assign it to a variable with the identifier "p1";
	- print out p1
	- print out just the field fName for p1
-----------------------------  */
type person struct {
	fName string
	lName string
}

func main() {
	p1 := person{"Todd", "McLeod"}
	fmt.Println(p1)
	fmt.Println(p1.fName)
}
