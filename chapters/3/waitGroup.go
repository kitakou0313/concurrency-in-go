package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello() {
	fmt.Println("Hello from Goroutine!")
	defer wg.Done()
}

func main() {
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
