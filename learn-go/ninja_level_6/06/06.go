package main

import (
	"fmt"
)

/* HANDS ON # 6 -----------------------------------
1. build and use an anonymous func
----------------------------------- */

func main() {
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}
