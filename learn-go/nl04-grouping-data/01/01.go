package main

import (
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. using a COMPOSITE LITERAL:
	- create an ARRAY which holds 5 VALUES of TYPE int
	- assign VALUES to each index position.
2. range over the array and print the values out using format printing
	- print out the TYPE of the array
----------------------------------- */
func main() {
	arr := [5]int{5, 9, 10, 13, 24}

	fmt.Println("Values in 'arr':")
	for i, v := range arr {
		fmt.Printf(" arr[%v]: %v\n", i, v)
	}

	fmt.Printf("\nType of 'arr':\n %T\n", arr)
}
