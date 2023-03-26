package main

import (
	"fmt"
)

const size int = 5

func largestInt(input [size]int) int {
	maxInt := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > maxInt {
			maxInt = input[i]
		}
	}
	return maxInt
}

func smallestInt(input [size]int) int {
	minInt := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < minInt {
			minInt = input[i]
		}
	}
	return minInt
}

func main() {
	var arr = [size]int{3, 1, 1, 50, 24}
	fmt.Println(smallestInt(arr))
	fmt.Println(largestInt(arr))
}
