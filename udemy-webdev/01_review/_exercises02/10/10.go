package main

import (
	"fmt"
)

/* QUESTION 10 -----------------------------
	- create a new type called 'transportation'. The underlying type of this
		new type is interface. For this interface, any other type that has a
		method < transportationDevice() string > will implicitly implement
		this interface;
	- create a func called 'report' that takes a value of type 'transportation'
		as an argument.  The func should print the string returned by
		transportationDevice() being called;
	- call 'report' passing in a value of type truck;
	- call 'report' passing in a value of type sedan;
-----------------------------  */
type transportation interface {
	transportationDevice() string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("This truck is", t.color, "and has", t.doors, "doors.")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("This sedan is", s.color, "and has", s.doors, "doors.")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	t := truck{vehicle{2, "black"}, true}
	s := sedan{vehicle{4, "red"}, false}

	report(t)
	report(s)
}
