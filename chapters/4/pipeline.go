package main

import "fmt"

func main() {
	// パイプラインで行う複数のステージを切り出すことで，各ステージの懸念事項や値などを切り出せる
	// いくつかのステージを作成する
	multiply := func(values []int, multiplier int) []int {
		multipliedValue := make([]int, len(values))
		for i, v := range values {
			multipliedValue[i] = v * multiplier
		}
		return multipliedValue
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add((multiply(ints, 2)), 1) {
		fmt.Println(v)
	}
}
