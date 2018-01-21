package main

import (
	"fmt"
)

/* QUESTION 5 -----------------------------
	- take the struct 'person' from the previous exercise and add a field called
		"favFood" that stores a slice of string;
	- for the variable 'p1', use a composite literal to add values to the field favFood;
	- print out the values in favFood;
	- also print out the values for 'favFood' by using a for range loop
-----------------------------  */
type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	p1 := person{
		"Nilson",
		"Molina",
		[]string{"pizza", "bannanas", "mashed potatoes"},
	}
	fmt.Println("favFoods:", p1.favFood)

	fmt.Println("\nListing favFoods with 'for range loop':")
	for _, val := range p1.favFood {
		fmt.Println("-", val)
	}
}
