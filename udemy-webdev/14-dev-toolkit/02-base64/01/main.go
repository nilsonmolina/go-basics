package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "hello"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Printf("STRING:\n%v\nLength: %v\n\n", s, len(s))
	fmt.Printf("ENCODED (base64):\n%v\nLength: %v\n\n", s64, len(s64))
}
