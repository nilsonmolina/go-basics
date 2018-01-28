package main

import "fmt"

/* HANDS ON # 5 -----------------------------------
1. show the comma ok idiom starting with this code.

func main() {
	c := make(chan int)

	v, ok :=
	fmt.Println(v, ok)

	close(c)

	v, ok =
	fmt.Println(v, ok)
}
----------------------------------- */

func main() {
	c := make(chan int)

	go func() {
		c <- 13
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
