package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* HANDS ON # 1 -----------------------------------
1. in addition to the main goroutine, launch two additional goroutines
	- each additional goroutine should print something out
2. use waitgroups to make sure each goroutine finishes before your program exists
----------------------------------- */

var wg sync.WaitGroup

func main() {
	fmt.Println("------ Initial State ------")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("---- Start Go Routines ----")

	wg.Add(2)
	go foo()
	go bar()

	wg.Wait()

	fmt.Println("----- End Go Routines -----")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 4; i++ {
		runtime.Gosched()
		fmt.Println("foo:", i, "\tgoroutine:", runtime.NumGoroutine())
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 4; i++ {
		runtime.Gosched()
		fmt.Println("bar:", i, "\tgoroutine:", runtime.NumGoroutine())
	}
	wg.Done()
}
