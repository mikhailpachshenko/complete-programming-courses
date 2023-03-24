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
	arrI := []int{1, 15, 3, 12, 9, 20}
	arrII := []int{2, 12, 9, 18, 0, 12, 22, 31}
	fmt.Printf("arr: %T, min: %d, max: %d.\n", arrI, minimum(arrI), maximum(arrI))
	fmt.Printf("arr: %T, min: %d, max: %d.\n", arrII, minimum(arrII), maximum(arrII))
}
