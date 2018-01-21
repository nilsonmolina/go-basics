package main

import (
	"fmt"
)

/* QUESTION 2 -----------------------------
	- initialize a 'MAP' using a composite literal where
		the key is a string and the value is an int
	- print out the map;
	- range over the map printing out just the key;
	- range over the map printing out both the key and value;
-----------------------------  */
func main() {
	friendMap := map[string]int{
		"Nilson":  26,
		"Edouard": 19,
		"Mark":    22,
		"Linh":    25,
	}
	fmt.Println("map:", friendMap)

	fmt.Println("\nEvery key in map:")
	for key := range friendMap {
		fmt.Println("Name:", key)
	}

	fmt.Println("\nEvery key/value pair in map:")
	for key, val := range friendMap {
		fmt.Println("Name:", key, "	Age:", val)
	}
}
