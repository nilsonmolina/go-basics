package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct {
	Name, Address, City, Zip, Region string
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "Hotel California",
			Address: "13876 Mobster Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "Southern",
		},
		hotel{
			Name:    "Nostromo Dorms",
			Address: "42 Sunset Boulevard",
			City:    "Fremont",
			Zip:     "94555",
			Region:  "Central",
		},
		hotel{
			Name:    "Vineyard Estates",
			Address: "4387",
			City:    "Tahoe",
			Zip:     "95612",
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
