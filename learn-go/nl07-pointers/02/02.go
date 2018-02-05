package main

import (
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
1. create a person struct
2. create a func called “changeMe” with a *person as a parameter
	- change the value stored at the *person address
	* to dereference a struct, use (*value).field
		- p1.first
		- (*p1).first
	* as an exception, if the type of x is a named pointer type and (*x).f
	is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.
		- https://golang.org/ref/spec#Selectors
3. create a value of type person
	- print out the value
4. in func main call “changeMe”
5. in func main print out the value
----------------------------------- */

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Todd McLeod"
	// (*p).name = "Nilson Molina" // also works
}

func main() {
	p := person{
		name: "James Bond",
	}
	fmt.Println(p)

	changeMe(&p)
	fmt.Println(p)
}
