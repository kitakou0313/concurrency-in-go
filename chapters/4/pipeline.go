package main

func main() {

	multiply := func(values []int, multiplier int) []int {
		multipliedValue := make([]int, len(values))
		for i, v := range values {
			multipliedValue[i] = v * multiplier
		}
		return multipliedValue
	}
}
