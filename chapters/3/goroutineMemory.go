package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memoConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutine = 1e4
	wg.Add(numGoroutine)

	before := memoConsumed()
	for i := 0; i < numGoroutine; i++ {
		go noop()
	}
	after := memoConsumed()

	fmt.Printf("Before:%.3f kb\n", before)
	fmt.Printf("After:%.3f kb\n", after)
}
