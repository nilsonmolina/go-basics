package main

import (
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
1. Print every rune code point of the uppercase alphabet three times.
Your output should look like this:
	65
		U+0041 'A'
		U+0041 'A'
		U+0041 'A'
	66
		U+0042 'B'
		U+0042 'B'
		U+0042 'B'
... through the rest of the alphabet characters
----------------------------------- */

func main() {

	for letter := 'A'; letter <= 'Z'; letter++ {
		fmt.Println(letter)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", letter) // fmt.Printf("\t%U '%c'\n", letter, letter)
		}
	}
}
