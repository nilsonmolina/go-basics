package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			fmt.Printf("--- %s CONN CLOSED ---\n", conn.RemoteAddr())
			break
		}
		i++
	}

	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(conn)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(conn)
	case rMethod == "POST" && rURI == "/apply":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}
}

func handleIndex(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET Index</title>
	</head>
	<body>
		<h1>GET Index</h1>
		<a href="/">index</a><br />
		<a href="/apply">apply</a><br />
		<a href="/bad">bad url</a>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET Apply</title>
	</head>
	<body>
		<h1>GET Apply</h1>
		<a href="/">index</a><br />
		<a href="/apply">apply</a><br />
		<a href="/bad">bad url</a>
		<form action="/apply" method="post">
		<input type="hidden" value="In my good death">
		<input type="submit" value="submit">
		</form>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApplyPost(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>POST Apply</title>
	</head>
	<body>
		<h1>POST Apply</h1>
		<a href="/">index</a><br />
		<a href="/apply">apply</a><br />
		<a href="/bad">bad url</a>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleDefault(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET Default</title>
	</head>
	<body>
		<h1>GET Default</h1>
		<a href="/">index</a><br />
		<a href="/apply">apply</a><br />
		<a href="/bad">bad url</a>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
