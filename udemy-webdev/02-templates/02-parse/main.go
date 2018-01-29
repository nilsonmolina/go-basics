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

// init values only needed for func initParsedGlob()
func init() {
	initTpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	file := "tpl.gohtml"
	files := []string{
		"templates/one.gohtml",
		"templates/two.gohtml",
		"templates/vespa.gohtml",
	}
	glob := "templates/*.gohtml"

	for {
		menu()
		input, err := readInput()
		if err != nil {
			continue
		}

		if input == 1 {
			writeToStdOut(file)
		} else if input == 2 {
			writeToFile(file)
		} else if input == 3 {
			parseFiles(files)
		} else if input == 4 {
			parseGlob(glob, files)
		} else if input == 5 {
			initParseGlob(files)
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

func writeToStdOut(file string) {
	tpl, err := template.ParseFiles(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeToFile(file string) {
	tpl, err := template.ParseFiles(file)
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

func parseFiles(files []string) {
	tpl, err := template.ParseFiles(files[0])
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles(files[1], files[2])
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[2], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[1], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[0], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseGlob(glob string, files []string) {
	tpl, err := template.ParseGlob(glob)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[2], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[1], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[0], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func initParseGlob(files []string) {
	err := initTpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = initTpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[2], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = initTpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[1], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = initTpl.ExecuteTemplate(os.Stdout, strings.TrimPrefix(files[0], "templates/"), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
