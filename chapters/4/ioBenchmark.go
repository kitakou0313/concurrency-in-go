package main

import (
	"bufio"
	"io"
	"log"
	"os"
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

func tmpFileOrFatal() *os.File {
	file, err := os.CreateTemp("", "tmp")
	if err != nil {
		log.Fatal("error", err)
	}
	return file
}

func BenchmarkUnBuffer(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferWrite(b *testing.B) {
	befferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(befferedFile))
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for benchMark := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{benchMark.(byte)})
	}
}
