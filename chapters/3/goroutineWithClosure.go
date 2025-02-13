package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// salutaion := "Hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutaion = "Hello in goroutine"
	// }()
	// wg.Wait()
	// fmt.Println(salutaion)

	for _, salutaion := range []string{"Hello", "Greetings", "Good day"} {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(salutaion)
		}()
	}
	wg.Wait()
}
