package main

import (
	"fmt"
)

/* HANDS ON # 3 -----------------------------------
1. Using the code from the previous example, use SLICING to create
the following new slices which are then printed:
	- [42 43 44 45 46]
	- [47 48 49 50 51]
	- [44 45 46 47 48]
	- [43 44 45 46 47]
----------------------------------- */
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("x\t= ", x)
	fmt.Println("x[:5]\t= ", x[:5])
	fmt.Println("x[5:]\t= ", x[5:])
	fmt.Println("x[2:7]\t= ", x[2:7])
	fmt.Println("x[1:6]\t= ", x[1:6])
}
