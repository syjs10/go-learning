package main

import "fmt"

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }
	y = appendInt(x, 1, 2, 3, 4, 5)
	fmt.Printf("cap=%d\t%v\n", cap(y), y)
}
