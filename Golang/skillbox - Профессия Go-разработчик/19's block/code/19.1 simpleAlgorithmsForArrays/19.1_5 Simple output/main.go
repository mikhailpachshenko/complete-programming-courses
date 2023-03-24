package main

import (
	"fmt"
)

func minimum(input []int) int {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func maximum(input []int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func main() {
	arr := []int{5, 3, 3, 6, 12, 54, 27}
	arrII := []int{7, 12, 11, 18, 55, 17, 33, 99}
	fmt.Printf("min: %d, max: %d\n", minimum(arr), maximum(arr))
	fmt.Printf("min: %d, max: %d\n", minimum(arrII), maximum(arrII))
}
