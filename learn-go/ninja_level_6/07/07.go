package main

import (
	"fmt"
)

/* HANDS ON # 7 -----------------------------------
1. assign a func to a variable, then call that func
----------------------------------- */

func main() {
	f := func() {
		fmt.Println("This function is a variable.")
	}

	f()
	fmt.Printf("type: %T\n", f)
}
