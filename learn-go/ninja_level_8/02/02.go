package main

import (
	"encoding/json"
	"fmt"
)

/* HANDS ON # 2 -----------------------------------
1. Starting  with this code , unmarshal the JSON into a Go data structure.
	* Hint: https://mholt.github.io/json-to-go/

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

}
----------------------------------- */

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	// your code goes here
	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println("\nERRROR:", err)
	}

	fmt.Println("\nPeople:\n-------------------")
	for _, person := range people {
		fmt.Println("Name:\t", person.First, person.Last)
		fmt.Println("Age:\t", person.Age)
		fmt.Println("Sayings:")
		for _, saying := range person.Sayings {
			fmt.Println("   -", saying)
		}
		fmt.Print("\n")
	}
}
