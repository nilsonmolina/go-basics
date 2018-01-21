package main

import (
	"fmt"
	"math"
)

// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

/*
Capitalization Matters!
	- 'shape' used in package
	- 'Shape' used outside of package (exported)
	* if exported, a comment is required above the declaration
*/
type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	// fmt.Printf("The area of the %T is: %f\n", s, s.area())
	fmt.Println("The area of the shape is: ", s.area())
}

func main() {
	sq := square{10}
	circ := circle{6}
	info(sq)
	info(circ)
}
