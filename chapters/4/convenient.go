package main

import "fmt"

func main() {
	// 与えられた配列の要素を無限に繰り返す
	repeat := func(
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
	take := func(
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

	done := make(chan interface{})
	for v := range take(done, repeatFn(done, func() interface{} { return "a" }), 3) {
		fmt.Println(v)
	}
}
