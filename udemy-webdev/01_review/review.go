package main

import (
	"fmt"
)

// Composition
type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says, \"How are you this morning, James.\"")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, "says, \"Shaken, but not stirred.\"")
}

// Polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// Slice
	fmt.Println("---------------")
	fmt.Println("-    SLICE    -")
	fmt.Println("---------------")

	xi := []int{2, 4, 7, 9, 42}
	fmt.Println("slice 'xi' =", xi)

	y := make([]int, 0, 100)
	fmt.Println("slice 'y' =", y)
	y = append(y, 8)
	y = append(y, 13, 42)
	fmt.Println("slice 'y' =", y)

	// Map
	fmt.Println("\n---------------")
	fmt.Println("-     MAP     -")
	fmt.Println("---------------")

	m := map[string]int{
		"Todd":    45,
		"Nina":    25,
		"Patrick": 27,
	}
	fmt.Println("map data: ", m)
	fmt.Printf("map type: %T\n", m)
	fmt.Println("Listed using for:")
	for key, value := range m {
		fmt.Println(key, "-", value)
	}
	fmt.Println("Listed using for w/out value variable:")
	for key := range m {
		fmt.Println(key, "-", m[key])
	}
	fmt.Println("map[\"Nina\"]:", m["Nina"])

	// Composition
	fmt.Println("\n---------------")
	fmt.Println("- COMPOSITION -")
	fmt.Println("---------------")

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	p1.speak()

	agent007 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	agent007.speak()
	// agent007.person.speak()

	// Polymorphism
	fmt.Println("\n----------------")
	fmt.Println("- POLYMORPHISM -")
	fmt.Println("----------------")

	saySomething(p1)
	saySomething(agent007)
}
