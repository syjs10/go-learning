package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("%x, %x", c1, c2)
}
