package main

import (
	"fmt"
)

/* HANDS ON # 10 -----------------------------------
1. Write down what these print:
	- fmt.Println( true && true )
	- fmt.Println( true && false )
	- fmt.Println( true || true )
	- fmt.Println( true || false )
	- fmt.Println( !true )
----------------------------------- */

func main() {
	fmt.Println("Required:")
	fmt.Println("true && true:\t", true && 1 == 1)
	fmt.Println("true && false:\t", true && false)
	fmt.Println("true || false:\t", true || false)
	fmt.Println("!true:\t\t", !true)
	fmt.Println("true || true:\t", true || 1 == 1)
}
