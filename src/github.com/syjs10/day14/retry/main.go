package main

import (
	"fmt"
	"time"
)

func main() {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		tmp := time.Second << uint(tries)
		fmt.Printf("retrying....  %d\n", tmp)
		time.Sleep(tmp)
	}
}
