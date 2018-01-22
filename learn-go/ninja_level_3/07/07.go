package main

import (
	"fmt"
)

/* HANDS ON # 7 -----------------------------------
1. Building on the previous hands-on exercise, create a program that
uses “else if” and “else”.
----------------------------------- */

func main() {
	x := 8
	if x == 7 || x == 13 {
		fmt.Println("Lucky number", x)
	} else if x < 0 {
		fmt.Println("Why so negative?")
	} else {
		fmt.Println("Sorry,", x, "is not a lucky number")
	}
}
