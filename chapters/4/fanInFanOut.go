package main

import (
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	// 与えられた関数を実行し，その結果を送信するチャンネルを返す
	repeatFn := func(
		done <-chan interface{},
		fun func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)

			for {
				select {
				case <-done:
					return
				case valueStream <- fun():
				}
			}
		}()
		return valueStream
	}

	toInt := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int, 0)
		defer close(intStream)

		go func() {
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()

		return intStream
	}

	// 500000までの整数をランダムに返す関数
	rand := func() interface{} {
		return int(math.Floor(rand.Float64() * 500000))
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
}
