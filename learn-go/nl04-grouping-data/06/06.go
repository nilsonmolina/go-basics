package main

import (
	"fmt"
)

/* HANDS ON # 6 -----------------------------------
1. Create a slice to store the names of all of the states in the
United States of America. What is the length of your slice? What
is the capacity? Print out all of the values, along with their
index position in the slice, without using the range clause. Here
is a list of the states:
	` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
	` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
	` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
	` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`,
	` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
	` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`,
	` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
	` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
	` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
	` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
----------------------------------- */
func main() {
	// states := make([]string, 50, 50)
	states := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`,
		` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
		` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`,
		` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
		` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
		` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}

	fmt.Println("'states' length:\t", len(states))
	fmt.Println("'states' capacity:\t", cap(states))

	fmt.Println("'states' values:")
	for i := 0; i < len(states); i++ {
		fmt.Printf("\tstates[%v]:\t%v\n", i, states[i])
	}
}
