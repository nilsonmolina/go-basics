package main

import (
	"fmt"
)

/* HANDS ON # 4 -----------------------------------
1. create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
3. in func main
	a. print out the value of the variable “x”
	b. print out the type of the variable “x”
	c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	d. print out the value of the variable “x”
----------------------------------- */

type hotdog int

var x hotdog

func main() {
	fmt.Println("Value of 'x' =", x)
	fmt.Printf("Type of 'x'  = %T\n", x)

	x = 42
	fmt.Println("\n'x' now equals:", x)
}
