package main

import (
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
1. using a COMPOSITE LITERAL:
	- create a SLICE of TYPE int
	- assign 10 VALUES
2. range over the slice and print the values out using format printing
	- print out the TYPE of the slice
----------------------------------- */
func main() {
	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Println("Values in 'slice':")
	for i, v := range slice {
		fmt.Printf(" slice[%v]: %v\n", i, v)
	}

	fmt.Printf("\nType of 'slice':\n %T\n", slice)
}
