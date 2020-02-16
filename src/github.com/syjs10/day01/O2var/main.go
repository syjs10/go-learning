package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "æµ‹è¯•"
	age = 16
	isOk = true

	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name: %s\n", name)
	fmt.Println(age)

	var s1 string = "s1"
	fmt.Println(s1)
	var s2 = "s1"
	fmt.Println(s2)
	s3 := "s3"
	fmt.Println(s3)
}
