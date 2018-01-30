package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/* INSTRUCTIONS -----------------------------------
ON MAC
1. start main.go (go run main.go), then go to a web browser and enter this:
	localhost:8080
----------------------------------- */

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never make it here, because scanner stays scanning and listening forever
	io.WriteString(conn, "\nCode here\n")
}
