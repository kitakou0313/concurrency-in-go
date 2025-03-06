package main

import (
	"io"
	"testing"
)

func repeat(
	done <-chan interface{},
	values ...interface{},
) <-chan interface{} {
	valuesStream := make(chan interface{})

	go func() {
		defer close(valuesStream)

		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valuesStream <- v:
				}
			}
		}
	}()

	return valuesStream
}

// 与えられたチャンネルから最初のn個を取得する
func take(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

func performWrite(b testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for benchMark := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{benchMark.(byte)})
	}
}

func Benchmark() {

}

func main() {

}
