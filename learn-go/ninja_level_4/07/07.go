package main

import (
	"fmt"
)

/* HANDS ON # 7 -----------------------------------
1. create a slice of a slice of string ([][]string). Store the following
data in the multi-dimensional slice:
	- "James", "Bond", "Shaken, not stirred"
	- "Miss", "Moneypenny", "Helloooooo, James."
2. range over the records, then range over the data in each record.
----------------------------------- */
func main() {
	multi := [][]string{
		{"James", "Bond", "Shaken, not stirred."},
		{"Miss", "Moneypenny", "Hellooooo, James."},
	}

	for i, outerV := range multi {
		fmt.Printf("multi[%v]: %v\n", i, outerV)
		for j, innerV := range multi[i] {
			fmt.Println("  index:", j, "\tvalue:", innerV)
		}
	}
}
