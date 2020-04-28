package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 5, 6}
	fmt.Println("p == ", p[1:4])

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d, cap=%d, %v\n", s, len(x), cap(x), x)
}
