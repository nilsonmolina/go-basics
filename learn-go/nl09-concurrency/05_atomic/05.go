package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/* HANDS ON # 5 -----------------------------------
1. fix the race condition you created in exercise #4 by using package atomic
----------------------------------- */

func main() {
	var wg sync.WaitGroup
	var incrementer int64

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
			atomic.AddInt64(&incrementer, 1)
			fmt.Println("func", i, "incrementer =", atomic.LoadInt64(&incrementer))
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("----- End Go Routines -----")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("Incrementer:\t", incrementer)
}
