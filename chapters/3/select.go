package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0

loop:
	for {
		select {
		case <-done:
			break loop
		default:
			workCounter++
			time.Sleep(time.Second)
		}
	}

	fmt.Printf("%v \n", workCounter)
}
