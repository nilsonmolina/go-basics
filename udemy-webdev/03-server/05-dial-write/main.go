package main

import (
	"fmt"
	"net"
)

/* INSTRUCTIONS -----------------------------------
1. start tcp server from 02-tcp-read
2. start this program
----------------------------------- */

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you!")
}
