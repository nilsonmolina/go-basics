package main

import (
	"fmt"
)

/* QUESTION 16-----------------------------
	ADDING ONTO PREVIOUS QUESTION
	- create another type called 'flamingo'. Make the underlying type bool;
	- give 'flamingo' a method with this signature: greeting() and have
		it print "Hello, I am pink and beautiful and wonderful.";
	- create a new 'swampCreature'. The underlying type is interface. The
		behavior whiche the swampCreature interface defines is such that
		any type which has the greeting() method will implicitly implement it;
	- create a func called 'bayou' which takes a value of swampCreature as an
		argument. Have this func print out the greeting for whatever swampCreature
		might be passed in;
	- create a value of type flamingo;
	- call the bayou function twice, once passing in the gator and another with flamingo
-----------------------------  */

type gator int
type flamingo bool

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator.")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func main() {
	var g1 gator
	bayou(g1)

	var f1 flamingo
	bayou(f1)
}
