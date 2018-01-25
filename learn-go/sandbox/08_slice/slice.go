package main

import (
	"fmt"
)

func main() {
	// APPEND
	fmt.Println("---------------")
	fmt.Println("-    APPEND   -")
	fmt.Println("---------------")
	x := []int{1, 2, 3, 4}
	fmt.Println("x:\t\t", x)

	x = append(x, 11, 12, 13, 14)
	fmt.Println("x + n1, n2, etc:", x)

	y := []int{21, 22, 23, 24}
	x = append(x, y...)
	fmt.Println("x + y:\t\t", x)

	z := []int{31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	x = append(x, z[:4]...)
	fmt.Println("x + z:\t\t", x)

	// DELETING
	fmt.Println("\n---------------")
	fmt.Println("-    DELETE   -")
	fmt.Println("---------------")
	fmt.Println("x:\t\t", x)

	// you can remove with an append
	x = append(x[:8], x[12:]...)
	fmt.Println("remove 20's:\t", x)

	// MAKE
	fmt.Println("\n---------------")
	fmt.Println("-     MAKE    -")
	fmt.Println("---------------")

	// make([]T, length, capacity)
	x1 := make([]int, 3, 5)
	fmt.Println("x1:", x1, "\n length:", len(x1), "| capacity:", cap(x1))

	x1 = append(x1, 1, 1)
	fmt.Println("\nx1:", x1, "\n length:", len(x1), "| capacity:", cap(x1))

	x1 = append(x1, 2)
	fmt.Println("\nx1:", x1, "\n length:", len(x1), "| capacity:", cap(x1))

	// MULTI-DIMENSIONAL
	fmt.Println("\n---------------------")
	fmt.Println("- MULTI-DIMENSIONAL -")
	fmt.Println("---------------------")
	n := []string{"Pizza", "Hotdog", "Steak", "Mashed Potatoes"}
	fmt.Println("1d array n:\t", n)

	m := []string{"Nilson", "Linh", "Mark", "Edouard"}
	fmt.Println("1d array m:\t", m)

	nm := [][]string{n, m}
	fmt.Println("2d array nm:\t", nm)

	// UNDERLYING ARRAY
	fmt.Println("\n--------------------")
	fmt.Println("- UNDERLYING ARRAY -")
	fmt.Println("--------------------")
	x2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("x2 := []int{0, 1, 2, 3, 4, 5}\t-", x2)

	y2 := append(x2[:2], x2[5:]...)
	fmt.Println("y2 := append(x2[:2], x2[5:]...)\t-", y2)

	fmt.Println("\nIf capacity of slice is large enough, append re-uses the underlying array.")
	fmt.Println("x2 = ", x2)
	fmt.Println("y2 = ", y2)
	fmt.Println("\t* Notice that x2 has changed!")

}
