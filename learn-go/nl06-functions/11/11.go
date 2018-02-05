package main

import (
	"fmt"
)

/* HANDS ON # 11 -----------------------------------
PERSONAL HANDS ON ABOUT RECURSION (WAS SUPPOSE TO MAKE A VID)
1. write a func that solves factorials with recursion
2. write a func that solves factorials with iteration
----------------------------------- */

func recursiveFactorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

func iterativeFactorial(n int) int {
	total := 1
	for ; n > 1; n-- {
		total *= n
	}
	return total
}

func main() {
	fmt.Println("Recursive Factorial of 5:", recursiveFactorial(5))
	fmt.Println("Iterative Factorial of 5:", iterativeFactorial(5))
}
