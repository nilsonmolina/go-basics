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

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			fmt.Fprintf(conn, "  - %s\n", ln)
		}

		// We never make it here, because the scanner keeps on reading.
		fmt.Println("Code got here.")
		io.WriteString(conn, "\nI see you connected.")

		conn.Close()
	}
}
