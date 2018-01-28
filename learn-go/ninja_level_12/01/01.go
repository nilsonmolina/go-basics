package main

import (
	"fmt"

	"github.com/nilsonmolina/go-basics/learn-go/ninja_level_12/01/dog"
)

/* HANDS ON # 1 -----------------------------------
1. Create a dog package. The dog package should have an exported func “Years”
which takes human years and turns them into dog years (1 human year = 7 dog years).
Document your code with comments. Use this code in func main.
	- run your program and make sure it works
	- run a local server with godoc and look at your documentation.


// SOLUTION
1. in terminal, run the command:
	godoc -http=:8080
2. navigate to documentation/packages/dog to see my documentation
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
}
