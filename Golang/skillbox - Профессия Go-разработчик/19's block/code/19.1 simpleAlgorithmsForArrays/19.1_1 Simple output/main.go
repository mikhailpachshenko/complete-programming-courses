package main

import (
	"fmt"
)

const size = 5

func maximum(input [size]int) int {
	maximum := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > maximum {
			maximum = input[i]
		}
	}
	return maximum
}

func minimum(input [size]int) int {
	minimum := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < minimum {
			minimum = input[i]
		}
	}
	return minimum
}

func main() {
	arr := [size]int{10, 49, 48, 20, 39}
	fmt.Println(maximum(arr))
	fmt.Println(minimum(arr))
}
