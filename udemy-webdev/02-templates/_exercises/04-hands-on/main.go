package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type meal struct {
	Name        string
	Description string
	Price       float64
}

type menu struct {
	Type  string
	Meals []meal
}

type restaurant struct {
	Name string
	Menu []menu
}

func main() {
	restaurants := []restaurant{
		restaurant{
			Name: "Mel's Diner",
			Menu: []menu{
				menu{
					Type: "Breakfast",
					Meals: []meal{
						meal{
							Name:        "Oatmeal",
							Description: "yum yum",
							Price:       4.95,
						},
					},
				},
				menu{
					Type: "Lunch",
					Meals: []meal{
						meal{
							Name:        "Hamburger",
							Description: "Delicous good eating for you",
							Price:       6.95,
						},
					},
				},
				menu{
					Type: "Dinner",
					Meals: []meal{
						meal{
							Name:        "Pasta Bolognese",
							Description: "From Italy delicious eating",
							Price:       7.95,
						},
					},
				},
			},
		},
		restaurant{
			Name: "Cracker Barrel",
			Menu: []menu{
				menu{
					Type: "Breakfast",
					Meals: []meal{
						meal{
							Name:        "Juice Orange",
							Description: "Delicious drinking in throat squeezed fresh",
							Price:       2.95,
						},
					},
				},
				menu{
					Type: "Lunch",
					Meals: []meal{
						meal{
							Name:        "French Fries",
							Description: "French eat potatoe fingers",
							Price:       2.95,
						},
					},
				},
				menu{
					Type: "Dinner",
					Meals: []meal{
						meal{
							Name:        "Bistro Potatoe",
							Description: "Bistro bar wood American bacon",
							Price:       6.95,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
