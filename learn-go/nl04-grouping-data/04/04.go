package main

import (
	"fmt"
)

/* HANDS ON # 4 -----------------------------------
1. start with this slice
	- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	- append to that slice this value: 52
	- print out the slice
2. in ONE STATEMENT append to that slice these values: 53, 54, 55
	- print out the slice
3. append to the slice this slice
	- y := []int{56, 57, 58, 59, 60}
	- print out the slice
----------------------------------- */
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("x:\n", x)

	x = append(x, 52)
	fmt.Println("\nappend 52:\n", x)

	x = append(x, 53, 54, 55)
	fmt.Println("\nappend 53, 54, 55:\n", x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("\nappend []int{56, 57, 58, 59, 60}:\n", x)
}
