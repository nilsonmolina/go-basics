package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* ------------------------------
RUNNING THIS FILE
	- To run this file, install Go and set the environment variable using the following command:
		export PATH=$PATH:/usr/local/go/bin

	- Then in terminal, navigate to this directory and use either of the following commands:
		go build && go ./intro
			or
		go run herewe.go
	* http server will be running and listening on localhost:8080
--------------------------------- */

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
	for value := range favNums2 {
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
		- a collection of key-value pairs
		- you can have maps within maps
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-	  MAPS	      -")
	fmt.Println("-----------------------")

	presAge := make(map[string]int)

	presAge["Theodore Roosevelt"] = 42
	presAge["John F. Kennedy"] = 38
	fmt.Println("Key:	Theodore Roosevelt\nValue: ", presAge["Theodore Roosevelt"])
	fmt.Println("len(presAge): ", len(presAge))

	delete(presAge, "John F. Kennedy")
	fmt.Println("len(presAge): ", len(presAge))

	/* ------------------------------
	FUNCTIONS
		- you can return more than one value with functions
		- variadic functions are functions that accept an unknown number of values
		- Go functions may be closures. A closure is a function value that
			references variables from outside its body.
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-	FUNCTIONS     -")
	fmt.Println("-----------------------")

	// These functions are defined outside of main. (below)
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum :", addThemUp(listNums))

	number1, number2 := next2Values(5)
	fmt.Println(number1, number2)

	// Variadic function
	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// Closures
	number3 := 3
	doubleNum := func() int {
		number3 *= 3
		return number3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	/* ------------------------------
	RECURSION
		- calling a function from inside itself
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-	RECURSION     -")
	fmt.Println("-----------------------")

	// The function is defined outside of main (below)
	inputNum := 3
	fmt.Println("factorial of: ", inputNum, "\n", factorial(inputNum))

	/* ------------------------------
	DEFER & PANIC
		- Defer is used to ensure that a function call is performed later
			in a program’s execution, usually for purposes of cleanup
		- A panic typically means something went unexpectedly wrong. Mostly
			we use it to fail fast on errors that shouldn’t occur during
			normal operation, or that we aren’t prepared to handle gracefully.
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-    DEFER & PANIC    -")
	fmt.Println("-----------------------")

	// // printTwo() will not run until main() completes
	// defer printTwo()
	// printOne()

	// functions defined outside of main (below)
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))
	demPanic()

	/* ------------------------------
	POINTERS
		- Pointers allow you to pass memory addresses and assign values in a
			different way.
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-      POINTERS     -")
	fmt.Println("-----------------------")

	x := 0
	fmt.Println("x Address =", &x)
	fmt.Println("x =", x)
	changeXVal(x)
	fmt.Println("x =", x)
	changeXValNow(&x)
	fmt.Println("x =", x)

	yPtr := new(int)
	fmt.Println("yPtr Address =", yPtr)
	fmt.Println("yPtr Value = ", *yPtr)
	changeXValNow(yPtr)
	fmt.Println("yPtr Value = ", *yPtr)

	/* ------------------------------
	STRUCTS
		- Structs allow us to define our own data types and are useful for grouping
			data together.
		- Go automatically handles conversion between values and pointers for method
			calls. You may want to use a pointer receiver type to avoid copying on
			method calls or to allow the method to mutate the receiving struct.
	--------------------------------- */
	fmt.Println("-----------------------")
	println("- STRUCTS/INTERFACES -")
	fmt.Println("-----------------------")

	rect1 := Rectangle{0, 50, 10, 10}

	fmt.Println("Rectangle is", rect1.width, "units wide")
	fmt.Println("Area of rectangle = ", rect1.area())

	squar := Square{20, 50}
	circ := Circle{4}
	fmt.Println("Area of Square =", getArea(squar))
	fmt.Println("Area of Circle =", getArea(circ))

	/* ------------------------------
	Strings
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-       Strings       -")
	fmt.Println("-----------------------")

	sampString := "Hello World"
	fmt.Println("sampString:		", sampString)
	fmt.Println("Contains \"lo\":		", strings.Contains(sampString, "lo"))
	fmt.Println("Index \"lo\":		", strings.Index(sampString, "lo"))
	fmt.Println("Count \"l\":		", strings.Count(sampString, "l"))
	fmt.Println("Replace \"l\", \"x\", 2:	", strings.Replace(sampString, "l", "x", 2))

	csvString := "1,2,3,4,5,6"
	fmt.Println("\ncsvString:	", csvString)
	fmt.Println("split:		", strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "b"}
	fmt.Println("\nlistOfLetters:	", listOfLetters)
	sort.Strings(listOfLetters)
	fmt.Println("sorted:		", listOfLetters)
	joinedLetters := strings.Join(listOfLetters, ", ")
	fmt.Println("joined:		", joinedLetters)

	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println("\n", listOfNums)

	/* ------------------------------
	FILE IO
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-       FILE IO       -")
	fmt.Println("-----------------------")

	// Write file
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text put within a new file.")
	file.Close()

	//Read File
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)

	/* ------------------------------
	CASTING
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-       CASTING       -")
	fmt.Println("-----------------------")

	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	fmt.Println("randInt:	", randInt, "	to float64:	", float64(randInt))
	fmt.Println("randFloat:	", randFloat, "	to int:		", int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println("randString:	", randString, "	to int:		", newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println("randString2:	", randString2, "	to float64:	", newFloat)

	/* ------------------------------
	HTTP SERVER
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-     HTTP SERVER     -")
	fmt.Println("-----------------------")

	fmt.Println("Uncomment the code in this section to activate http server")
	// // handler functions defined outside of main (below)
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/earth", handler2)

	// http.ListenAndServe(":8080", nil)

	/* ------------------------------
	GO ROUTINES
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-     GO ROUTINES     -")
	fmt.Println("-----------------------")

	for i := 0; i < 2; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 6000)

	/* ------------------------------
	CHANNELS
	--------------------------------- */
	fmt.Println("-----------------------")
	println("-    CHANNELS    -")
	fmt.Println("-----------------------")

	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		fmt.Println("--------------------")
		time.Sleep(time.Millisecond * 2000)
	}
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

func subtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}

// Used for recursion
func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1)
}

// Used for defer
func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func safeDiv(num1, num2 int) int {
	defer func() {
		// fmt.Println(recover())
		recover()
	}()

	solution := num1 / num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")
}

func changeXVal(x int) {
	x = 2
}

func changeXValNow(x *int) {
	*x = 2
}

// Rectangle : this is a rectangle struct
type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

// Shape : this is an interface
type Shape interface {
	area() float64
}

// Square : this is a Square struct
type Square struct {
	height float64
	width  float64
}

// Circle : this is a Circle struct
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.height * s.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

func count(id int) {
	for i := 0; i < 6; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough and Send for Sauce")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 100)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add Sauce and Send", pizza, "for toppings")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 100)
}

func addToppings(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond * 100)
}
