package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go sayHello()
}
