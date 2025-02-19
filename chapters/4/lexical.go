package main

import "fmt"

// 読み込み専用のチャンネルを返す -> 書き込みの権限をchanOwner関数にのみ拘束する
var chanOwner = func() <-chan int {
	results := make(chan int, 5)
	go func() {
		defer close(results)
		for i := 0; i < 5; i++ {
			results <- i
		}
	}()

	return results
}

var consumer = func(results <-chan int) {
	for result := range results {
		fmt.Printf("Receive: %d \n", result)
	}
}

func main() {
	results := chanOwner()
	consumer(results)
}
