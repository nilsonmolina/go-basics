package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

var tplFunction *template.Template
var fm = template.FuncMap{
	"toUpper":    strings.ToUpper,
	"firstThree": firstThree,
	"fmtTime":    formatTime,
	"fdouble":    double,
	"fsquare":    square,
	"fsqroot":    sqRoot,
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	tplFunction = template.Must(template.New("").Funcs(fm).ParseFiles("templates/func/08-functions.gohtml"))
}

func main() {
	for {
		menu()
		input, err := readInput()
		if err != nil {
			continue
		}

		if input == 1 {
			passToBasic()
		} else if input == 2 {
			assignToVariable()
		} else if input == 3 {
			passSlice()
		} else if input == 4 {
			passMap()
		} else if input == 5 {
			passStruct()
		} else if input == 6 {
			passSliceStruct()
		} else if input == 7 {
			passComposite()
		} else if input == 8 {
			passFunction()
		} else if input == 9 {
			usingMethods()
		} else if input == 0 {
			break
		}

		fmt.Println()
	}
}

func menu() {
	fmt.Println("----------------------------")
	fmt.Println("Passing Data into Templates")
	fmt.Println("----------------------------")
	fmt.Println("1. Pass data to Basic Template (01-Data.gohtml)")
	fmt.Println("2. Assign to Variable (02-variables.gohtml)")
	fmt.Println("3. Pass a Slice (03-slice.gohtml)")
	fmt.Println("4. Pass a Map (04-map.gohtml)")
	fmt.Println("5. Pass a Struct (05-struct.gohtml)")
	fmt.Println("6. Pass a Slice of Struct (06-slice-struct.gohtml)")
	fmt.Println("7. Pass a Struct of Structs (07-composite.gohtml)")
	fmt.Println("8. Pass a Function (func/08-functions.gohtml)")
	fmt.Println("9. Using Methods (09-methods.gohtml)")
	fmt.Println("0. QUIT")
}

func readInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	rs, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.Trim(rs, "\n"))
}

func passToBasic() {
	// I passed 13 into the template, so the {{.}} in the 01-data.gohtml file became 13
	err := tpl.ExecuteTemplate(os.Stdout, "01-data.gohtml", 13)
	if err != nil {
		log.Fatalln(err)
	}
}

func assignToVariable() {
	// As shown in the 02-variables.gohtml file, you can assign the '.' to a variable. The
	// data passed ("The number 13") is then assigned to that variable and can be called on.
	err := tpl.ExecuteTemplate(os.Stdout, "02-variables.gohtml", "The number 13")
	if err != nil {
		log.Fatalln(err)
	}
}

func passSlice() {
	games := []string{"Halo", "Mass Effect", "The Last of Us",
		"Uncharted", "Final Fantasy", "Super Smash Bros."}

	// If only one template, I can just use Execute w/out giving template name
	err := tpl.ExecuteTemplate(os.Stdout, "03-slice.gohtml", games)
	if err != nil {
		log.Fatalln(err)
	}
}

func passMap() {
	m := map[string]int{
		"Nilson":  26,
		"Linh":    25,
		"Mark":    22,
		"Edouard": 19,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "04-map.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}
}

type person struct {
	Name  string
	Age   int
	Admin bool
}

func passStruct() {
	// Make sure to use 'exported' fields or it won't work - (capitalize field names)
	p := person{
		Name: "Nilson",
		Age:  26,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "05-struct.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}

func passSliceStruct() {
	p1 := person{Name: "Nilson", Age: 26}
	p2 := person{Name: "Linh", Age: 25}
	p3 := person{Name: "Mark", Age: 22}
	p4 := person{Name: "Edouard", Age: 19}

	people := []person{p1, p2, p3, p4}

	err := tpl.ExecuteTemplate(os.Stdout, "06-slice-struct.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
}

type vehicle struct {
	Manufacturer string
	Model        string
	Doors        int
}

func passComposite() {
	p1 := person{Name: "Nilson", Age: 26}
	p2 := person{Name: "Linh", Age: 25}
	p3 := person{Name: "Mark", Age: 22}
	p4 := person{Name: "Edouard", Age: 19}
	people := []person{p1, p2, p3, p4}

	v1 := vehicle{Manufacturer: "Ford", Model: "F150", Doors: 2}
	v2 := vehicle{Manufacturer: "Tesla", Model: "Model 3", Doors: 4}
	v3 := vehicle{Manufacturer: "Honda", Model: "Civic Si", Doors: 4}
	cars := []vehicle{v1, v2, v3}

	data := struct {
		People    []person
		Transport []vehicle
	}{
		People:    people,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "07-composite.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return string([]rune(s)[:3]) // previous code: return s[:3]
}
func formatTime(t time.Time) string {
	return t.Format("01-02-2006")
}
func double(x int) int {
	return x + x
}
func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func passFunction() {
	p1 := person{Name: "Nilson", Age: 26, Admin: true}
	p2 := person{Name: "Linh", Age: 25, Admin: true}
	p3 := person{Name: "Mark", Age: 22}
	p4 := person{Name: "Edouard", Age: 19}
	p5 := person{Name: "Priscilla", Age: 24}
	people := []person{p1, p2, p3, p4, p5}

	data := struct {
		People []person
		Time   time.Time
		Value  int
	}{
		People: people,
		Time:   time.Now(),
		Value:  3,
	}

	// If only one template, I can just use Execute w/out giving template name
	err := tplFunction.ExecuteTemplate(os.Stdout, "08-functions.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func (p person) GetDogAge() int {
	return p.Age * 7
}

func (p person) GetBirthYear(age int) int {
	return time.Now().Year() - age
}

func usingMethods() {
	p := person{
		Name: "Linh Lam",
		Age:  24,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "09-methods.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}

}
