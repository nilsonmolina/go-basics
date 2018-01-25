package main

import "fmt"

const (
	_      = iota
	kbIota = 1 << (iota * 10)
	mbIota = 1 << (iota * 10)
	gbIota = 1 << (iota * 10)
)

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Println("\nUsing regular multiplication:")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Println("\nUsing iota:")
	fmt.Printf("%d\t\t\t%b\n", kbIota, kbIota)
	fmt.Printf("%d\t\t\t%b\n", mbIota, mbIota)
	fmt.Printf("%d\t\t%b\n", gbIota, gbIota)
}
