package main

import (
	"fmt"
)

/* QUESTION 17 -----------------------------
	- using the short declaration operator, create a variable with the
		identifier 's' and assign "I'm sorry Dave, I can't do that" to it;
	- print s;
	- print s converted to a slice of byte;
	- print s converted to a slice of byte and then converted back to a string;
	- using slicing, print just "I'm sorry Dave";
	- using slicing, print just "Dave I can't";
	- using slicing, print just "can't do that";
	- print every letter of s with one rune (character) on each line
-----------------------------  */

func main() {
	s := "I'm sorry Dave, I can't do that"
	fmt.Println("Original string:\n", s)

	fmt.Println("\nConverted to slice of byte:\n", []byte(s))
	fmt.Println("\nConverted to slice and back to string:\n", string([]byte(s)))
	fmt.Println("\nUsing slicing s[:14]\n", s[:14])
	fmt.Println("\nUsing slicing s[:14]\n", s[10:23])
	fmt.Println("\nUsing slicing s[:14]\n", s[18:])

	fmt.Println("\nOne rune per line:")
	for _, char := range s {
		fmt.Println(string(char))
	}
}
