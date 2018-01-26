package main

import (
	"fmt"
)

/* HANDS ON # 3 -----------------------------------
1. use the “defer” keyword to show that a deferred func runs
after the func containing it exits.
----------------------------------- */

func main() {
	defer fmt.Println("the end of the function...")
	fmt.Println("Welcome to...")
}
