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
	arrI := []int{5, 12, 8, 11, 8, 9, 55}
	arrII := []int{4, 11, 12, 11, 8, 9, 7}
	fmt.Printf("arr: %T, min: %d, max: %d\n", arrI, minimum(arrI), maximum(arrI))
	fmt.Printf("arr: %T, min: %d, max: %d\n", arrII, minimum(arrII), maximum(arrII))
}
