package main

import (
	"fmt"
)

const size int = 5

func minimum(input [size]int) int {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func maximum(input [5]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func main() {
	var arr = [size]int{5, 2, 1, 12, 7}
	fmt.Println(minimum(arr))
	fmt.Println(maximum(arr))
}
