package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("Rakib_Sultan\n"))
	fmt.Printf("%x", h.Sum(nil))
}
