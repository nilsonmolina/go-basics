package main

import (
	"fmt"
)

/* QUESTION 6 -----------------------------
	- Add a method to type 'person' with the identifier "walk". Have the
		func return this string: <persons first name> is walking.
	- To return a string, use fmt.Sprintln. Call the method assigning the
		returned string to a variable with the identifier "s". Print out s.
	* Remember, you make a func a method by giving the func a receiver:
		func (r receiver) identifier(parameters) (returns) { <code /> }
-----------------------------  */
type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking.")
}
func main() {
	p := person{"Todd", "McLeod", []string{"bananas", "strawberries", "apples"}}
	s := p.walk()
	fmt.Println(s)
}
