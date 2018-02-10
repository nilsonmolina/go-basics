package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	s := "test@example.com"
	c := getCode(s)
	fmt.Printf("%v\t%v\n", s, c)

	c = getCode(s)
	fmt.Printf("%v\t%v\n", s, c)

	s = "test@exampl.com"
	c = getCode(s)
	fmt.Printf("%v\t\t%v\n", s, c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
