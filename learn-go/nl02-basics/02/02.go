package main

import "fmt"

/* HANDS ON # 2 -----------------------------------
1. Using the following operators, write expressions and assign their values to variables:
	- ==
	- <=
	- >=
	- !=
	- <
	- >
Now print each of the variables.
----------------------------------- */

func main() {
	a := 13 == 42
	b := 13 <= 42
	c := 13 >= 42
	d := 13 != 42
	e := 13 < 42
	f := 13 > 42

	fmt.Println("13 == 42 :", a)
	fmt.Println("13 <= 42 :", b)
	fmt.Println("13 >= 42 :", c)
	fmt.Println("13 != 42 :", d)
	fmt.Println("13 < 42  :", e)
	fmt.Println("13 > 42  :", f)
}
