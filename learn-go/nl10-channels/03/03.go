package main

import "fmt"

/* HANDS ON # 3 -----------------------------------
1. starting with this code, pull the values off the channel using a for range loop

func main() {
	c := gen()
	receive(c)
	fmt.Println("about to exit")
}
func gen() <-chan int {
	c := make(chan int)
	for i := 0; i < 100; i++ {
		c <- i
	}
	return c
}
----------------------------------- */

func main() {
	c := gen()
	receive(c)
	fmt.Println("about to exit")
}
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
