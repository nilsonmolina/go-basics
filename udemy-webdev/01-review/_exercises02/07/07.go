package main

/* QUESTION 7 -----------------------------
	- create a new type 'vehicle'. The underlying type is a struct. The
		fields: doors, color;
	- create two new types: 'truck' and 'sedan'. The underlying type is a struct;
	- embed the 'vehicle' type in both 'truck' and 'sedan';
	- give 'truck' the field "fourWheel" which will be set to bool.
	- give 'sedan' the field "luxury" which will be set to bool.
-----------------------------  */

type vehicle struct {
	doors string
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

func main() {}
