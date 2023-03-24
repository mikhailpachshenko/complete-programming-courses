package main

import (
	"fmt"
)

const size int = 5

func mininum(input [size]int) int {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func maximum(input [size]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func main() {
	var arr = [size]int{3, 5, 9, 12}
	fmt.Println(mininum(arr))
	fmt.Println(maximum(arr))

	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arr, arr, arr)
}
