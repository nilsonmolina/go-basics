package main

import (
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. create your own type “person” which will have an underlying type of “struct”
so that it can store the following data:
	- first name
	- last name
	- favorite ice cream flavors
2. create two VALUES of TYPE person. Print out the values, ranging over the
elements in the slice which stores the favorite flavors.
----------------------------------- */

type person struct {
	fName       string
	lName       string
	favIceCream []string
}

func main() {
	p1 := person{
		fName: "Todd",
		lName: "Mcleod",
		favIceCream: []string{
			"Strawberry",
			"White Mocha",
			"Dulce de leche",
		},
	}
	p2 := person{
		fName: "Linh",
		lName: "Voyo",
		favIceCream: []string{
			"Coffee",
			"Greek Yogurt",
			"Snickers",
		},
	}

	fmt.Println("First:\t", p1.fName)
	fmt.Println("Last:\t", p1.lName)
	fmt.Println("Fav Ice Cream:")
	for i, v := range p1.favIceCream {
		fmt.Println(" ", i, v)
	}

	fmt.Println("\nFirst Name:\t", p2.fName)
	fmt.Println("Last Name:\t", p2.lName)
	fmt.Println("Fav Ice Cream:")
	for i, v := range p2.favIceCream {
		fmt.Println(" ", i, v)
	}
}
