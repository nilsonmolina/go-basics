package main

import (
	"fmt"
	"sort"
)

/* HANDS ON # 5 -----------------------------------
1. Starting  with this code,  sort the []user by:
	- age
	- last
2. Also sort each []string “Sayings” for each user
	- print everything in a way that is pleasant
----------------------------------- */

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	// your code goes here
	fmt.Println("UNSORTED:\n------------------")
	for _, u := range users {
		fmt.Printf("%v %v, %v\n", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Println("    -", s)
		}
		fmt.Println()
	}

	sort.Sort(byLast(users))
	fmt.Println("SORTED (by Last):\n------------------")
	for _, u := range users {
		fmt.Printf("%v %v, %v\n", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Println("    -", s)
		}
		fmt.Println()
	}

	sort.Sort(byAge(users))
	fmt.Println("SORTED (by Age):\n------------------")
	for _, u := range users {
		fmt.Printf("%v %v, %v\n", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Println("    -", s)
		}
		fmt.Println()
	}

	fmt.Println("SORTED (the 'Sayings'):\n------------------")
	for _, u := range users {
		fmt.Printf("%v %v, %v\n", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Println("    -", s)
		}
		fmt.Println()
	}
}

type byLast []user

func (a byLast) Len() int           { return len(a) }
func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

type byAge []user

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
