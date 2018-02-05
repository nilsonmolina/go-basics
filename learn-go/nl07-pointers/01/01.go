package main

import (
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. create a value and assign it to a variable.
2. print the address of that value.
----------------------------------- */

func main() {
	x := 13
	fmt.Println("Value of x:\t", x)
	fmt.Println("Address of x:\t", &x)
}
