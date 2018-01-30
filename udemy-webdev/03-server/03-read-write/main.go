package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	fmt.Fprintln(conn, "\n-------------------------------")
	fmt.Fprintln(conn, " Welcome to the TCP server 1.3 ")
	fmt.Fprintln(conn, "-------------------------------")
	fmt.Fprintln(conn, "My construct has been upgraded! I can now 'read' and 'write' from my connection.")
	fmt.Fprintln(conn, "I also have the ability to accept multiple connections concurrently, rather")
	fmt.Fprintln(conn, "than just one at a time! I will now repeat anything you say to me for a few")
	fmt.Fprintln(conn, "seconds and then this connection will be closed.")
	fmt.Fprintln(conn, "\nparrot_mode_engaged:\t\ttrue")
	fmt.Fprintln(conn, "conn_self_destruct_timer:\t30 seconds")
	fmt.Fprintln(conn, "awaiting user input:")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n\n", ln)
	}
	defer conn.Close()

	// we can now make it here, because connection has a deadline
	fmt.Println("Connection has been closed!")
}
