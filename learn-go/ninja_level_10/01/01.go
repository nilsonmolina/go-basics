package main

import "fmt"

/* HANDS ON # 1 -----------------------------------
USING THE FOLLOWING CODE:
1. get the code working:
	- using func literal, aka, anonymous self-executing func
	- using a buffered channel

func main() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
----------------------------------- */

func main() {
	// // DOES NOT WORK - channels block
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)

	// func literal solution
	c := make(chan int)
	go func() {
		c <- 13
	}()
	fmt.Println(<-c)

	// // buffered channel solution
	// c := make(chan int, 1)
	// c <- 42
	// fmt.Println(<-c)
}
