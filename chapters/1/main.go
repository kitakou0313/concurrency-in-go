package main

import "fmt"

func main() {
	var data int

	go func() {
		data++ // Write
	}()

	if data == 0 { // Read
		fmt.Println("The val is %v") // Read
	}
}
