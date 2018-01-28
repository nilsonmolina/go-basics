package main

import (
	"fmt"
)

/* QUESTION 14 -----------------------------
	ADDING ONTO PREVIOUS QUESTION
	- use CONVERSION to assign the value of g1 to x
	* Since g1 is of type gator, but it's underlying type is an int
		so we can use conversion to convert the value to an int.
	* CONVERSION is also known as CASTING in other languages.
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

	fmt.Println("\nWith Conversion, we can set x = g1")
	x = int(g1)
	fmt.Println("g1   =", g1)
	fmt.Println("x    =", x)
}
