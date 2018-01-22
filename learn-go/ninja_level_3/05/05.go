package main

import (
	"fmt"
)

/* HANDS ON # 5 -----------------------------------
1. Print out the remainder (modulus) which is found for each number
between 10 and 100 when it is divided by 4.
----------------------------------- */

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "% 4 =", i%4)
	}
}
