package main

import "fmt"

/* HANDS ON # 2 -----------------------------------
1. get this code running:

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
----------------------------------- */

func main() {
	//cs := make(chan<- int) // SOLUTION: changed receiver channel to a general channel
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
