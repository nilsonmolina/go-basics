package main

import (
	"fmt"
)

/* HANDS ON # 4 -----------------------------------
1. Create a user defined struct with the identifier “person”
	- the fields:
		* first * last * age
2. attach a method to type person with
	- the identifier “speak”
	- the method should have the person say their name and age
3. create a value of type person
4. call the method from the value of type person
----------------------------------- */

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	p := person{
		first: "John",
		last:  "Doe",
		age:   36,
	}
	p.speak()
}
