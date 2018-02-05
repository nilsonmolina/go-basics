package main

import (
	"fmt"
)

/* HANDS ON # 6 -----------------------------------
1. Using  iota, create 4 constants for the  NEXT  4 years.
2. Print the constant values.
----------------------------------- */

const (
	a = 2018 + iota
	b = 2018 + iota
	c = 2018 + iota
	d = 2018 + iota
)

func main() {
	fmt.Printf("a: %v\nb: %v\nc: %v\nd: %v\n", a, b, c, d)
}
