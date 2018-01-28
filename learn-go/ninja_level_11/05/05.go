package main

import "fmt"

/* HANDS ON # 5 -----------------------------------
1. We are going to learn about testing next. For this exercise,
take a moment and see how much you can figure out about testing by
reading the testing documentation & also Caleb Doxsey’s article
on testing. See if you can get a basic example of testing working.
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
