package main

import (
	"fmt"
	"time"
)

// 送信用，受信用のGo routineの定義は使われ方を表している
var doWork = func(
	done <-chan interface{},
	strings <-chan string,
) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited")
		defer close(terminated)

		for {
			select {
			case s := <-strings:
				// 別Go routineから受け取った処理
				fmt.Println(s)
			case <-done:
				return

			}
		}
	}()

	return terminated
}

func main() {
	done := make(chan interface{})

	// nilを渡しておりfor loopで止まるはず
	// 今回はdoneチャンネルを渡して1s後にCloseしているため，停止できた
	isDoworkGoRoutineTerminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine")
		close(done)
	}()

	<-isDoworkGoRoutineTerminated
	fmt.Println("Done")
}
