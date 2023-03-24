package main

import (
	"fmt"
)

const (
	size    = 5
	newSize = 2
)

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

func wrongCalculation(input [size]int) {
	for i := 0; i < 6; i++ {
		fmt.Println(input[i])
	}
}

func main() {
	arr := [size]int{10, 49, 48, 20, 39}
	fmt.Println(minimum(arr))
	fmt.Println(maximum(arr))

	wrongCalculation(arr)

	newArraysI := [...]int{10, 11, 12, 13}
	newArraysII := [...]int{10}

	smallestI := newArraysI[0]
	smallestII := newArraysII[1]

	if len(newArraysII) >= 2 {
		newSmallest := newArraysII[0]
		newSecondSmallest := newArraysII[1]
	} else {
		panic(newArraysII)
	}
}
