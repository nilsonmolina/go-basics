package main

import (
	"fmt"
)

/* HANDS ON # 6 -----------------------------------
1. write a program that
	- puts 100 numbers to a channel
	- pull the numbers off the channel and print them
----------------------------------- */

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}
}
