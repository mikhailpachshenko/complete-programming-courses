package main

import (
	"fmt"
)

const size = 5

func minimum(input [size]int) int {
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
	arr := [size]int{2, 1, 1, 15, 14}
	fmt.Println(minimum(arr))
	fmt.Println(maximum(arr))
}
