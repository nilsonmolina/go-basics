package main

import (
	"fmt"
)

/* QUESTION 8 -----------------------------
	- using the vehicle, truck, and sedan structs: using a composite literal,
		create a value of type truck and assign values to the fields;
	- using a composite literal, create a value of type sedan and assign
		values to the fields;
	- print each of these values;
	- print out a single field from each of these values;
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

func main() {
	t := truck{vehicle{2, "black"}, true}
	s := sedan{vehicle{4, "red"}, false}

	fmt.Println("Truck:", t)
	fmt.Println("Sedan:", s)

	fmt.Println("\nThe truck is", t.color, "and has", t.doors, "doors.")
	fmt.Println("The sedan is", s.color, "and has", s.doors, "doors.")
}
