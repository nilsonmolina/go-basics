package main

import (
	"fmt"
)

/* HANDS ON # 9 -----------------------------------
1. a “callback” is when we pass a func into a func as an argument.
For this exercise, pass a func into a func as an argument
----------------------------------- */

func loop(f func(n int) int) {
	for i := 0; i < 3; i++ {
		fmt.Println(f(12))
	}
}

func increment(n int) int {
	n++
	return n
}

func main() {
	loop(increment)
}
