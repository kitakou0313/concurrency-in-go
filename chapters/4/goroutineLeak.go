package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("dowork exited")
			defer close(completed)

			// 引数のチャンネルからの受信待ち
			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	// nilチャンネルを渡しているため，main go routineの完了まで生成したGo routineは生存し続ける
	doWork(nil)
	fmt.Println("Do some work")
}
