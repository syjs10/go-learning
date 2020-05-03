package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Jiang Song", 25}
	z := Person{"Zhang San", 100}
	fmt.Println(a, z)
}
