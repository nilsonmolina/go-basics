package main

import (
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. create a func with the identifier foo that returns an int
2. create a func with the identifier bar that returns an int and a string
3. call both funcs
4. print out their results
----------------------------------- */
func foo() int {
	return 42
}

func bar() (int, string) {
	return 13, "Lucky number"
}

func main() {
	fmt.Println("foo(): ", foo())

	x, s := bar()
	fmt.Println("bar(): ", x, s)
}
