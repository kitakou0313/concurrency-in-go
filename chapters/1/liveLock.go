package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Microsecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)

		atomic.AddInt32(dir, 1)

		takeStep()

		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprintf(out, ". Success")
			return true
		}

		takeStep()

		atomic.AddInt32(dir, -1)
		return false
	}
}
