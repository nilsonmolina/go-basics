package main

import "fmt"

/* HANDS ON # 5 -----------------------------------
USING THE CODE FROM THE PREVIOUS EXERCISE
1. at the package level scope, using the “var” keyword, create a VARIABLE with the
IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
2. in func main
	a. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	b. then use the “=” operator to ASSIGN that value to “y”
	c. print out the value stored in “y”
	d. print out the type of “y”
----------------------------------- */

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println("Value of 'x' =", x)
	fmt.Printf("Type of 'x'  = %T\n", x)
	x = 42
	fmt.Println("\n'x' now equals:", x)

	y = int(x) // Will not work like this: y = x
	fmt.Println("Value of 'y' =", y)
	fmt.Printf("Type of 'y'  = %T\n", y)
}
