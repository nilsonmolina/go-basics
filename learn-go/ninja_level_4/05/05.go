package main

import (
	"fmt"
)

/* HANDS ON # 5 -----------------------------------
To DELETE from a slice, we use APPEND along with SLICING.
 1. Starting with this slice:
 x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	- use APPEND & SLICING to get these values and ASSIGN to a
		variable “y” and then print: [42, 43, 44, 48, 49, 50, 51]
----------------------------------- */
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("x:\n", x)

	y := append(x[:3], x[6:]...)
	fmt.Println("y:\n", y)
}
