package main

import (
	"fmt"
	"runtime"
)

/* HANDS ON # 6 -----------------------------------
1. create a program that prints out your OS and ARCH. Use the following commands to run it
	- go run
	- go build
	- go install
----------------------------------- */

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("Architecture:\t", runtime.GOARCH)
}
