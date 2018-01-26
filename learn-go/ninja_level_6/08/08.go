package main

import (
	"fmt"
)

/* HANDS ON # 8 -----------------------------------
1. create a func which returns a func
	- assign the returned func to a variable
	- call the returned func
----------------------------------- */

func retFunc() func() string {
	return func() string {
		return "This string came from a func that returned a func and was assigned to a variable."
	}
}

func main() {
	f := retFunc()

	fmt.Println(f())
}
