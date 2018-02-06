package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	fmt.Fprintln(conn, "\n------------------------")
	fmt.Fprintln(conn, " Welcome to Exercise #5 ")
	fmt.Fprintln(conn, "------------------------")
	fmt.Fprintln(conn, "This tcp server can both READ and WRITE to a connection, but for now")
	fmt.Fprintln(conn, "it will simply parrot whatever input you provide.  Concurrency has now")
	fmt.Fprintln(conn, "been enabled and multiple connections can be established. To quit simply")
	fmt.Fprintln(conn, "return an empty line. (visit w/ browser - http://localhost:8080)")
	fmt.Fprintf(conn, "\nMessage:\t")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			fmt.Printf("--- %s CONN CLOSED ---\n", conn.RemoteAddr())
			break
		}
		fmt.Println(conn.RemoteAddr(), "-", ln)
		fmt.Fprintf(conn, "  - %s\n", ln)
		fmt.Fprintf(conn, "\nMessage:\t")
	}

	fmt.Println("Code got here.")                  // written at server
	io.WriteString(conn, "\nI see you connected.") // written at client

	conn.Close()
}
