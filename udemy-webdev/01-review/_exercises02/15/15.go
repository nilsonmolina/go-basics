package main

import (
	"fmt"
)

/* QUESTION 15 -----------------------------
	- now add a method to type gator with this signature:
		greeting() and have it print "Hello, I am a gator";
	- create a value of type gator;
	- call the greeting() method from that value;
-----------------------------  */

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator.")
}

func main() {
	var g1 gator
	g1.greeting()
}
