package main

import (
	"fmt"
)

/* QUESTION 13 -----------------------------
	ADDING ONTO PREVIOUS QUESTION
	- can you assign 'g1' to 'x'? Why or why not?
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

	fmt.Println("\nYou cannot directly assign 'g1' to 'x' because they are two different types!")
	//x = g1
	//fmt.Println(x)
}
