package main

import (
	"fmt"
)

/* QUESTION 12 -----------------------------
	ADDING ONTO PREVIOUS QUESTION
	- using var, declare an identifier 'x' as type int;
	- print out 'x';
	- print the type of 'x';
-----------------------------  */

type gator int

func main() {
	var g1 gator
	g1 = 13
	fmt.Println("g1   =", g1)
	fmt.Printf("type = %T\n", g1)

	var x int
	fmt.Println("\nx    =", x)
	fmt.Printf("type = %T\n", x)
}
