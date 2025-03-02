package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))

		go func() {
			defer close(intStream)
			for _, v := range integers {
				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}()

		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiply int,
	) <-chan int {
		multipliedStream := make(chan int)

		go func() {
			defer close(multipliedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- v * multiply:
				}
			}
		}()

		return multipliedStream
	}

	done := make(chan interface{}, 0)
	defer close(done)

	inteStream := generator(done, 1, 2, 3, 4, 5)
	pipeline := multiply(done, inteStream, 5)

	for v := range pipeline {
		fmt.Println(v)
	}

}
