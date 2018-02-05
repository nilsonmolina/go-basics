package main

import "fmt"

/* HANDS ON # 3 -----------------------------------
1. create a struct “customErr” which implements the builtin.error interface.
2. create a func “foo” that has a value of type error as a parameter.
3. create a value of type “customErr” and pass it into “foo”.

HINT:
type customErr struct {
}
func (ce customErr) Error() string {
}
func main() {
	c1 := customErr{}
	foo(c1)
}
func foo(e error) {
	fmt.Println(e)
}
----------------------------------- */

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "english hard bad",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Printf("foo:\n   %v\n", e.Error())

	// // if I want to access the 'info' field, I need to use assertion:
	// fmt.Printf("foo:\n   error2: %v\n", e.(customErr).info)
}
