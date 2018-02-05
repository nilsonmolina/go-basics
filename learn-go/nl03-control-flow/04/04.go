package main

import (
	"fmt"
	"time"
)

/* HANDS ON # 4 -----------------------------------
1. Create a for loop using this syntax
	- for {}
2. Have it print out the years you have been alive.
----------------------------------- */

func main() {
	year := 1991
	for {
		if year > time.Now().Year() {
			break
		}
		fmt.Println(year)
		year++
	}
}
