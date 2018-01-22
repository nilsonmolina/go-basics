package main

import (
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. using the short declaration operator,  ASSIGN  these  VALUES  to  VARIABLES
with the IDENTIFIERS “x” and “y” and “z”
	a. 42
	b. “James Bond”
	c. true
2. now print the values stored in those variables using
	a. a single print statement
	b. multiple print statements
----------------------------------- */

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("All variables:\n", x, y, z)

	fmt.Println("\nIndividual variables:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
}
