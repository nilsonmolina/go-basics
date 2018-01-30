package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"strings"
)
import text "text/template"

type page struct {
	Title   string
	Heading string
	Input   string
}

var textTpl *text.Template
var htmlTpl *template.Template

func init() {
	textTpl = text.Must(text.ParseFiles("tpl.gohtml"))
	htmlTpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	for {
		menu()
		input, err := readInput()
		if err != nil {
			continue
		}

		if input == 1 {
			textTemplate()
		} else if input == 2 {
			htmlTemplate()
		} else if input == 0 {
			break
		}

		fmt.Println()
	}
}

func menu() {
	fmt.Println("---------------------------------")
	fmt.Println("Text Templates VS HTML Templates")
	fmt.Println("---------------------------------")
	fmt.Println("1. XSS using Text Template")
	fmt.Println("2. XSS using HTML Template")
	fmt.Println("0. QUIT")
}

func readInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	rs, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.Trim(rs, "\n"))
}

func textTemplate() {
	home := page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := textTpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}

}

func htmlTemplate() {
	home := page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := htmlTpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}

}
