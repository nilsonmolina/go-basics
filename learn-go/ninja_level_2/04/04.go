package main

import "fmt"

/* HANDS ON # 4 -----------------------------------
Write a program that
	- assigns an int to a variable
	- prints that int in decimal, binary, and hex
	- shifts the bits  of that int over 1 position to the left, and
		assigns that to a variable
	- prints that variable in decimal, binary, and hex
----------------------------------- */

func main() {
	n := 13
	fmt.Printf("Decimal:\t%d\nBinary:\t\t%b\nHex:\t\t%#x\n", n, n, n)

	s := n << 1
	fmt.Printf("\nDecimal:\t%d\nBinary:\t\t%b\nHex:\t\t%#x\n", s, s, s)
}
