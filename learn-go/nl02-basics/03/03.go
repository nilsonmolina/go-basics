package main

import "fmt"

/* HANDS ON # 3 -----------------------------------
1. Create TYPED and UNTYPED constants.
2. Print the values of the constants.
----------------------------------- */

const (
	u     = 13
	t int = 42
)

func main() {
	fmt.Println("Untyped Const:\t", u, "\nTyped Const:\t", t)
}
