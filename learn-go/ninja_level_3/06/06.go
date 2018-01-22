package main

import (
	"fmt"
)

/* HANDS ON # 6 -----------------------------------
1. Create a program that shows the “if statement” in action.
----------------------------------- */

func main() {
	x := 13
	if x == 7 || x == 13 {
		fmt.Println("Lucky number", x)
	}
}
