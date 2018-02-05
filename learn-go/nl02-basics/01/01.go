package main

import (
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. Write a program that prints a number in:
	- decimal
	- binary
	- hex
----------------------------------- */

func main() {
	n := 42
	fmt.Printf("Decimal:\t%d\nBinary:\t\t%b\nHex:\t\t%#x\n", n, n, n)
}
