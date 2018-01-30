package main

import (
	"bufio"
	"fmt"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	fmt.Fprintln(conn, "\n---------------------")
	fmt.Fprintln(conn, "   rot13 server")
	fmt.Fprintln(conn, "---------------------")
	fmt.Fprintln(conn, "Welcome to the rot13 encryption engine (aka. The Caesar Cipher).")
	fmt.Fprintln(conn, "This is a simple tcp server in which you provide a message and we")
	fmt.Fprintln(conn, "will encrypt that message for you.  To unencrypt, simply provide the")
	fmt.Fprintln(conn, "encrypted message. Message will be converted to lowercase for simplicity.")
	fmt.Fprintf(conn, "\nMessage:\t")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		r := rot13([]byte(ln))

		fmt.Fprintf(conn, "Encrypted:\t%s\n", r)
		fmt.Fprintf(conn, "\nMessage:\t")
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v < 'a' || v > 'z' {
			r13[i] = v
		} else if v <= 'm' {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
