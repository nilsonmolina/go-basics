package main

import (
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
USING CODE FROM PREVIOUS QUESTION
1. store the values of type person in a map with the key of last name.
	- access each value in the map.
	- print out the values, ranging over the slice.
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

	people := map[string]person{
		p1.lName: p1,
		p2.lName: p2,
	}

	for _, v := range people {
		fmt.Printf("%v %v\n", v.fName, v.lName)
		fmt.Println("Favorite Ice Cream:")
		for i, flavor := range v.favIceCream {
			fmt.Println("", i, flavor)
		}
		fmt.Println("------------------")
	}

}
