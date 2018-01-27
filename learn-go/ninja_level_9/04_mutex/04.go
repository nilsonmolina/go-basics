package main

import (
	"fmt"
	"runtime"
	"sync"
)

/* HANDS ON # 4 -----------------------------------
1. Fix the race condition you created in the previous exercise by using a mutex
	* it makes sense to remove runtime.Gosched()

----------------------------------- */

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	incrementer := 0
	routines := 100

	fmt.Println("------ Initial State ------")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("Incrementer:\t", incrementer)
	fmt.Println("---- Start Go Routines ----")

	// wg.Add(routines) // Added it in function, but same result if done here
	for i := 0; i < routines; i++ {
		go func(i int) {
			wg.Add(1)
			m.Lock()
			x := incrementer
			runtime.Gosched()
			x++
			incrementer = x
			fmt.Println("func", i, "incrementer =", incrementer)
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("----- End Go Routines -----")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("Incrementer:\t", incrementer)
}
