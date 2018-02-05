package main

import "fmt"

/* HANDS ON # 8 -----------------------------------
1. Create a program that uses a switch statement with no switch expression specified.
----------------------------------- */

func main() {
	x := 13
	switch {
	case x == 13 || x == 7:
		fmt.Println("Lucky number", x)
	case x < 0:
		fmt.Println("Why so negative?")
	default:
		fmt.Println("Sorry,", x, "is not a lucky number")
	}
}
