package main

import "fmt"

func add(a, b int, pipe chan int) {
	var sum int
	sum = a + b
	pipe <- sum
}

func main() {
	var sum int
	pipe := make(chan int, 1)
	go add(1, 2, pipe)
	sum = <-pipe
	fmt.Println(sum)
}
