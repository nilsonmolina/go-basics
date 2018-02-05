package main

import (
	"fmt"
)

/* HANDS ON # 10 -----------------------------------
1. closure is when we have “enclosed” the scope of a variable in some
code block. For this hands-on exercise, create a func which “encloses”
the scope of a variable:
----------------------------------- */

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	f := incrementer()
	fmt.Println("f():", f())
	fmt.Println("f():", f())
	fmt.Println("f():", f())
	fmt.Println("f():", f())
	g := incrementer()
	fmt.Println("g():", g())
	fmt.Println("g():", g())
	fmt.Println("f():", f())
	fmt.Println("g():", g())
}
