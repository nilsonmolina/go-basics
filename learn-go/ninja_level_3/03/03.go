package main

import (
	"fmt"
	"time"
)

/* HANDS ON # 3 -----------------------------------
1. Create a for loop using this syntax
	- for condition { }
2. Have it print out the years you have been alive.
----------------------------------- */

func main() {
	year := 1991
	for year <= time.Now().Year() {
		fmt.Println(year)
		year++
	}
}
