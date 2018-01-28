package main

import (
	"fmt"
)

// create an interface type that both person and secretAgent implement
// declare a func with a parameter of the interfaces's type
// call that func in main and pass in a value of type person
// call that func in main and pass in a value of type secretAgent

type human interface {
	speak()
}

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says, \"How are you this morning, James.\"")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, "says, \"Shaken, but not stirred.\"")
}

func speak(h human) {
	h.speak()
}

func main() {
	p := person{"Todd", "McLeod"}
	speak(p)

	sa := secretAgent{person{"James", "Bond"}, true}
	speak(sa)
}
