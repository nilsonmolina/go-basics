package main

import (
	"fmt"
)

/* HANDS ON # 3 -----------------------------------
1. create a new type: vehicle. The underlying type is a struct.
	The fields:
		- doors
		- color
2. create two new types: truck & sedan. The underlying types is a struct.
	- embed the “vehicle” type in both truck & sedan.
	- give truck the field “fourWheel” which will be set to bool.
	- give sedan the field “luxury” which will be set to bool. solution
3. using the vehicle, truck, and sedan structs:
	- using a composite literal, create a value of type truck and assign values to the fields;
	- using a composite literal, create a value of type sedan and assign values to the fields.
4. print out each of these values.
5. print out a single field from each of these values.
----------------------------------- */

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// Best Practice to show fields and values
	t := truck{
		vehicle:   vehicle{2, "Space Grey"},
		fourWheel: true,
	}
	// But you can omit them to be more concise
	s := sedan{
		vehicle{4, "White"},
		false,
	}

	fmt.Println(t)
	fmt.Printf("This %v door truck is %v\n\n", t.doors, t.color)
	fmt.Println(s)
	fmt.Printf("This %v door sedan is %v\n\n", s.doors, s.color)
}
