package main

import (
	"fmt"
)

/* QUESTION 9 -----------------------------
	- give a method to both the 'truck' and 'sedan' types with the
		following signature: transportationDevice() string
	- have each func return a string saying what they do;
	- create a value of type truck and populate the fields;
	= create a value of type sedan and populate the fields;
	- call the method for each value;
-----------------------------  */

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

func main() {
	t := truck{vehicle{2, "black"}, true}
	s := sedan{vehicle{4, "red"}, false}

	fmt.Println(t.transportationDevice())
	fmt.Println(s.transportationDevice())
}
