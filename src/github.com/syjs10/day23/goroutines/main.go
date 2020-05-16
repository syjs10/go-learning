package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("1")
	say("0")
}
