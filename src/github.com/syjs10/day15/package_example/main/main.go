package main

import (
	"github.com/syjs10/day15/package_example/calc"
	"fmt"
)

func main() {
	sum := calc.Add(300, 500)
	sub := calc.Sub(500, 300)
	fmt.Println(sum)
	fmt.Println(sub)

}

