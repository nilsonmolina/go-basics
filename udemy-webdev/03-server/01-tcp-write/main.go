package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

/* INSTRUCTIONS -----------------------------------
ON MAC
1. start main.go (go run main.go), then go to another terminal window and enter this:
	Before macOS 10.13.1 (High Sierra)
		telnet localhost 8080
	After macOS 10.13.1 (High Sierra)
		nc localhost 8080
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

		io.WriteString(conn, "\n-------------------------------\n")
		fmt.Fprintln(conn, " Welcome to the TCP server 1.1 ")
		fmt.Fprintf(conn, "%v\n", "-------------------------------")
		fmt.Fprintln(conn, "My sole function is to simply listen and accept any connections to this tcp port.")
		fmt.Fprintln(conn, "This revision only has the ability to 'write' to the connection, but not read. My")
		fmt.Fprintln(conn, "creator is fairly new to web programming and hasn't figured out how to do anything")
		fmt.Fprintln(conn, "else. So yup... this is pretty much it................ it's so quiet here...")
		fmt.Fprintln(conn)

		conn.Close()
	}
}
