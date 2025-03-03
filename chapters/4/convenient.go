package main

func main() {
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
}
