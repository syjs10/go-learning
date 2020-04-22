package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	score float64
	next  *Student
}

func main() {
	var head Student
	head.Name  = "hua"
	head.Age   = 20
	head.Score = 100

	var sut1 Student
	sut1.Name  = "stu1"
	sut1.Age   = 23
	sut1.Score = 90

	head.next  = &sut1
	sut1.next  = nil

	var p *Student = &head

	for p.next != nil {
		fmt.Println(*p)
		p = p.next
	}
}

