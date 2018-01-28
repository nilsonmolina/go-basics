package main

import (
	"fmt"

	"github.com/nilsonmolina/go-basics/learn-go/ninja_level_13/02/quote"
	"github.com/nilsonmolina/go-basics/learn-go/ninja_level_13/02/word"
)

/* HANDS ON # 2 -----------------------------------
1. start with this code. Get the code ready to BET on (add benchmarks, examples, tests)
however do not write an example for the func that returns a map; and only write a test for
it as an extra challenge. Add documentation to the code. Run the following in this order:
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

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
