package main

import "fmt"

const pi = 3.1415926

const (
	statusOk = 200
	notFound = 404
)
const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	a3
)
const (
	b1 = iota
	b2
	_
	b3
)
const (
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// fmt.Printf("n1: %d\n", n1)
	// fmt.Printf("n2: %d\n", n2)
	// fmt.Printf("n3: %d\n", n3)
	// fmt.Printf("a1: %d\n", a1)
	// fmt.Printf("a2: %d\n", a2)
	// fmt.Printf("a3: %d\n", a3)
	// fmt.Printf("b1: %d\n", b1)
	// fmt.Printf("b2: %d\n", b2)
	// fmt.Printf("b3: %d\n", b3)

	// math.MaxFloat32
	s1 := "123测试字符串"
	s2 := []rune(s1)

	fmt.Println(len(s2))
	for _, c := range s2 {
		fmt.Println(string(c))
	}
}
