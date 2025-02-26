package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result, 0)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{
					Error:    err,
					Response: resp,
				}

				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{
		"example[.]com", "hogehoge[.]com",
	}

	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			// 全体のエラーに基づいた処理を行う
			fmt.Println("Error")
			fmt.Println(result.Error)
			continue
		}
	}
}
