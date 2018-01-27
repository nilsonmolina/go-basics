package main

import (
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
This exercise will reinforce our understanding of method sets:
1. create a type person struct
2. attach a method speak to type person using a pointer receiver
	- *person
3. create a type human interface
	- to implicitly implement the interface, a human must have the speak method
4. create func “saySomething”
	- have it take in a human as a parameter
	- have it call the speak method
5. show the following in your code
	- you CAN pass a value of type *person into saySomething
	- you CANNOT pass a value of type person into saySomething
----------------------------------- */

type human interface {
	speak()
}

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old.")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Todd McLeod",
		age:  46,
	}
	// saySomething(p)	// Does not work
	saySomething(&p)
}
