package main

import (
	"fmt"
	"time"
)

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(<-c)
		}
	}()
	go func() {
		time.Sleep(time.Microsecond * 1000)
		quit <- 0
	}()
	fib(c, quit)
}
