package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

/* INSTRUCTIONS -----------------------------------
1. start tcp server from 01-tcp-write
2. start this program
----------------------------------- */

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
