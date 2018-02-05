package main

import (
	"fmt"

	"github.com/nilsonmolina/go-basics/learn-go/ninja_level_13/01/dog"
)

/* HANDS ON # 1 -----------------------------------
1. start with this code. Get the code ready to BET on the code
(add benchmarks, examples, tests). Run the following in this order:
	- tests
	- benchmarks
	- coverage
	- coverage shown in web browser
	- examples shown in documentation in a web browser

// SOLUTION
1. created main_test.go and made BET code
2. ran the following commands:
	- go test
	- go test -bench .
	- go test -cover
	- go test -coverprofile c.out
	- go tool cover -html=c.out
	- godoc -http=:8080
----------------------------------- */

type canine struct {
	name string
	age  int
}

func main() {
	d := canine{
		name: "Mr. Cow",
		age:  dog.Years(3),
	}
	fmt.Println(d.name, "is", d.age, "in dog years.")
	dog.YearsAlt(3)
}
