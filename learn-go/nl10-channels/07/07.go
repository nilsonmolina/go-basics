package main

import (
	"fmt"
)

/* HANDS ON # 7 -----------------------------------
1. write a program that
	- launches 10 goroutines
		* each goroutine adds 10 numbers to a channel
	- pull the numbers off the channel and print them
----------------------------------- */

func main() {
	routines := 10
	num := 10

	for i := 0; i < routines; i++ {
		go launch(num)
	}
}

func launch(num int) {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < num; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}
}
