package main

import (
	"fmt"
)

/* HANDS ON # 4 -----------------------------------
1. create and use an anonymous struct.
----------------------------------- */

func main() {
	s := struct {
		first    string
		friends  map[string]int
		favGames []string
	}{
		first:    "Nilson",
		friends:  map[string]int{"Linh": 25, "Mark": 22, "Edouard": 19},
		favGames: []string{"Halo", "Mass Effect", "The Last of Us", "Uncharted"},
	}
	fmt.Println(s)
}
