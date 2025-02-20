package main

import (
	"bytes"
	"fmt"
	"sync"
)

var printData = func(wg *sync.WaitGroup, data []byte) {
	defer wg.Done()

	var buff bytes.Buffer
	for _, b := range data {
		fmt.Fprintf(&buff, "%c", b)
	}
	fmt.Println(buff.String())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	// dataはprintDataの後のため，引数としてしか受けてれない
	// スライスの部分集合を渡しているため，競合を防止できる
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
