package main

import "fmt"

func main() {
	/* ------------------------------
	DECLARING VARIABLES
		- variables are statically typed, so they cannot be changed (ex: cannot go from int to float)
		* can use println for debugging because it does not need dependencies,
			but fmt.Println is built into stdlib and is used for production code.
	--------------------------------- */
	fmt.Println("-----------------------")
	println("- DECLARING VARIABLES -")
	fmt.Println("-----------------------")
	//var age int = 13
	//var favNum float64 = 13.31
	var age = 13
	var favNum = 13.31
	randNum := 1
	fmt.Println(age, favNum, randNum)

	// like in most languages, float arithmetic is not precise
	var num1 = 1.000
	var num99 = 0.9999
	fmt.Println(num1 - num99)

	// constants cannot be changed after creating them
	const pi = 3.14159
	fmt.Println(pi)

	// strings
	var myName = "Nilson Molina"
	fmt.Println(myName, len(myName))
	fmt.Println(myName + " is a beast!!!")

	// booleans
	var isOver24 = true
	fmt.Println("Is over 24? -", isOver24)

	// printf examples - include \n since printf does not do it by default
	fmt.Printf("%.2f", pi)                 //prints a float with a precision of 2 decimals
	fmt.Printf("%T \n", pi)                //prints the variable type
	fmt.Printf("%t \n", isOver24)          //prints a boolean
	fmt.Printf("print int:	%d \n", 100)    //prints an int
	fmt.Printf("print binary:	%b \n", 100) //prints an int in binary
	fmt.Printf("print hex:	%x \n", 100)    //prints an int in hexadecimal
	fmt.Printf("print char 64:	%c \n", 64) //prints the ascii character for an int value
	fmt.Printf("print exp:	%e \n", pi)     //prints exponential

	/* ------------------------------
	LOGICAL OPERATORS
		- there are three logical operators inside of go

	--------------------------------- */
	fmt.Println("-----------------------")
	println("-  LOGICAL OPERATORS  -")
	fmt.Println("-----------------------")

	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = 	", !true)

	/* ------------------------------
		LOOPS & CONDITIONS
	--------------------------------- */
	fmt.Println("-----------------------")
	println("- LOOPS & CONDITIONS  -")
	fmt.Println("-----------------------")

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	yourAge := 26
	if yourAge >= 16 {
		fmt.Println("You can drive.")
	} else if yourAge >= 18 {
		fmt.Println("You can vote.")
	} else {
		fmt.Println("You can go have fun.")
	}

	switch yourAge {
	case 16:
		fmt.Println("You can drive.")
	case 18:
		fmt.Println("You can vote.")
	default:
		fmt.Println("You can go have fun.")
	}

	/* ------------------------------
	ARRAYS & SLICES
		- a slice is like an array, but you don't define the size
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-   ARRAYS & SLICES   -")
	fmt.Println("-----------------------")

	var favNums [5]int
	favNums[0] = 13
	favNums[1] = 26
	favNums[2] = 39
	favNums[3] = 52
	favNums[4] = 65
	fmt.Println(favNums)

	favNums2 := [5]int{13, 26, 39, 52, 65}
	for i, value := range favNums2 {
		fmt.Println(value, i)
	}
	for _, value := range favNums2 {
		fmt.Println(value)
	}

	numSlice := []int{5, 4, 3, 2, 1}
	fmt.Println("numSlice = ", numSlice)
	fmt.Println("numSlice[3:5] = ", numSlice[3:5])
	fmt.Println("numSlice[3:4] = ", numSlice[3:4])
	fmt.Println("numSlice[0:3] = ", numSlice[0:3])
	fmt.Println("numSlice[:3] = ", numSlice[:3])
	fmt.Println("numSlice[3:] = ", numSlice[3:])

	numSlice2 := make([]int, 5, 10)
	fmt.Println("numSlice2 = ", numSlice2)
	copy(numSlice2, numSlice)
	fmt.Println("copy(numSlice2, numSlice) = ", numSlice2)
	numSlice2 = append(numSlice2, 0, -1)
	fmt.Println("append(numSlice2, 0, -1) = ", numSlice2)

	/* ------------------------------
	MAPS
		- a slice is like an array, but you don't define the size
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-   ARRAYS & SLICES   -")
	fmt.Println("-----------------------")
}
