package main

import (
	"fmt"
	"math"
)

func main() {
	o := 0666
	x := int64(0xdeadbeef)
	fmt.Printf("十进制： %d，八进制： %[1]o，带符号的八进制： %#[1]o\n", o)
	fmt.Printf("十进制： %d，十六进制： %[1]x ，带符号的十六进制： %#[1]x，带符号的大写十六进制： %#[1]X\n", x)
	ascii   := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
	for x := 0; x <= 8; x++ {
		fmt.Printf("int: %d, e^x:%8.3f\n", x, math.Exp(float64(x)))
	}
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
}
