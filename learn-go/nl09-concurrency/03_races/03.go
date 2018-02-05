package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* HANDS ON # 3 -----------------------------------
1. Using goroutines, create an incrementer program
	- have a variable to hold the incrementer value
	- launch a bunch of goroutines. Each goroutine should
		- read the incrementer value (store it in a new variable)
		- yield the processor with runtime.Gosched()
		- increment the new variable
		- write the value in the new variable back to the incrementer variable
2. use waitgroups to wait for all of your goroutines to finish
	- the above will create a race condition.
	- Prove that it is a race condition by using the -race flag
----------------------------------- */

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	routines := 100

	fmt.Println("------ Initial State ------")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("Incrementer:\t", incrementer)
	fmt.Println("---- Start Go Routines ----")

	// wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func(i int) {
			wg.Add(1)
			x := incrementer
			runtime.Gosched()
			x++
			incrementer = x
			fmt.Println("func", i, "incrementer =", incrementer)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("----- End Go Routines -----")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("Incrementer:\t", incrementer)
}
