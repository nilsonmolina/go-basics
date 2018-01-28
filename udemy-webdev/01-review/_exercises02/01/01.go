package main

import (
	"fmt"
)

/* QUESTION 1 -----------------------------
	- initialize a 'SLICE' of int using a composite literal;
	- print out the slice;
	- range over the slice printing out just the index;
	- range over the slice printing out both the index and the value
-----------------------------  */
func main() {
	slice := []int{7, 13, 23, 42}
	fmt.Println("slice:", slice)

	fmt.Println("\nslice range showing only index:")
	for i := range slice {
		fmt.Println(i)
	}

	fmt.Println("\nslice range showing both index and value:")
	for i, val := range slice {
		fmt.Printf("slice[%d]: %d\n", i, val)
	}
}
