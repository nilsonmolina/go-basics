package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// init values only needed for func initParsedGlob()
var initTpl *template.Template

func init() {
	initTpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	for {
		menu()
		input, err := readInput()
		if err != nil {
			continue
		}

		if input == 1 {
			writeToStdOut()
		} else if input == 2 {
			writeToFile()
		} else if input == 3 {
			parseFiles()
		} else if input == 4 {
			parseGlob()
		} else if input == 5 {
			initParseGlob()
		} else if input == 0 {
			break
		}

		fmt.Println()
	}
}

func menu() {
	fmt.Println("----------------------------")
	fmt.Println("Welcome to Manual Templating")
	fmt.Println("----------------------------")
	fmt.Println("1. Write to stdout")
	fmt.Println("2. Write to file")
	fmt.Println("3. Parse files")
	fmt.Println("4. Parse GLOB")
	fmt.Println("5. Initialized Parse GLOB")
	fmt.Println("0. QUIT")
}

func readInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	rs, _ := reader.ReadString('\n')
	return strconv.Atoi(strings.Trim(rs, "\n"))
}

func writeToStdOut() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeToFile() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseFiles() {
	tpl, err := template.ParseFiles("templates/one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("templates/two.gohtml", "templates/vespa.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseGlob() {
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// If templates are changed after program runs, this will not show changes
// until program restarts, since templates are already in memory
func initParseGlob() {
	err := initTpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = initTpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = initTpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = initTpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
