package main

import (
	"fmt"
	"math"
)

/* HANDS ON # 5 -----------------------------------
1. create a type  SQUARE
2. create a type  CIRCLE
3. attach a method to each that calculates  AREA  and returns it
	- circle area= π r 2
	- squarearea=L*W
4. create a type  SHAPE  that defines an interface as anything that has the  AREA  method
5. create a func  INFO  which takes type shape and then prints the area
6. create a value of type square
7. create a value of type circle
8. use func info to print the area of square
9. use func info to print the area of circle
----------------------------------- */

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
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area =", s.area())
}

func main() {
	s := square{
		side: 13,
	}
	c := circle{radius: 13.13}
	info(s)
	info(c)
}
