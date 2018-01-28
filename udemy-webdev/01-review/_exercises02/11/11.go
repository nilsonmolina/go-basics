package main

import (
	"fmt"
)

/* QUESTION 11 -----------------------------
	- create a new type called 'gator'. The underlying type is an int;
	- using var, declare an identifier 'g1' as type gator;
	- assign a value to 'g1';
	- print out 'g1';
	- print the type of 'g1'
-----------------------------  */

type gator int

func main() {
	var g1 gator
	g1 = 13
	fmt.Println("g1 =  ", g1)
	fmt.Printf("type = %T\n", g1)
}
