package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Printf("STRING:\n%v\nLength: %v\n\n", s, len(s))
	fmt.Printf("ENCODED (base64):\n%v\nLength: %v\n\n", s64, len(s64))
}
