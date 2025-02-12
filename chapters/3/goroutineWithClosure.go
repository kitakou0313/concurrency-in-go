package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	salutaion := "Hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutaion = "Hello in goroutine"
	}()
	wg.Wait()
	fmt.Println(salutaion)
}
