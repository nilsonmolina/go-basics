package main

import (
	"encoding/json"
	"fmt"
)

/* HANDS ON # 1 -----------------------------------
1. Starting  with this code, marshal the []user to JSON. There is a
little bit of a curve ball in this hands-on exercise - remember to ask
yourself what you need to do to EXPORT a value from a package.

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Moneypenny",
		age:   27,
	}

	u3 := user{
		first: "M",
		age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
}
----------------------------------- */

// User : this comment is needed to remove the warning the linter is throwing
type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(bs), err)

}
