package main

import (
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
1. create a func with the identifier foo that
	- takes in a variadic parameter of type int
	- pass in a value of type []int into your func (unfurl the []int)
	- returns the sum of all values of type int passed in
2. create a func with the identifier bar that
	- takes in a parameter of type []int
	- returns the sum of all values of type int passed in
----------------------------------- */

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	return foo(xi...)
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println("n:\t", n)
	fmt.Println("foo(n):\t", foo(n...))
	fmt.Println("bar(n):\t", bar(n))
}
