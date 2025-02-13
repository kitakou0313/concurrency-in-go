package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// salutaion := "Hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutaion = "Hello in goroutine"
	// }()
	// wg.Wait()
	// fmt.Println(salutaion)

	for _, salutaion := range []string{"Hello", "Greetings", "Good day"} {
		wg.Add(1)

		// 各Goroutineは同じメモリ空間で実行される
		// salutaionへの参照を全Goroutineが持っている状態
		// Goroutineの処理実行時にはループが終了しているので，salutaionの値はループの最後の値になっている
		go func() {
			defer wg.Done()
			fmt.Println(salutaion)
		}()
	}
	wg.Wait()
}
