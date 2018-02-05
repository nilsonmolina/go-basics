package main

import (
	"fmt"
)

/* HANDS ON # 5 -----------------------------------
1. Create a variable of type string using a raw string literal.
2. Print it.
----------------------------------- */

func main() {
	str := `
	This was created using 
	a raw string literal.
	`
	fmt.Println(str)
}
