package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		defer v2.mu.Lock()

		fmt.Printf("v1 = %v, v2 = %v", v1.value, v2.value)
	}

	var a, b value
	wg.Add(2)

	go printSum(&a, &b) // v1 -> v2でロック取得
	go printSum(&b, &a) // v2 -> v1でロック取得

	wg.Wait()
}
