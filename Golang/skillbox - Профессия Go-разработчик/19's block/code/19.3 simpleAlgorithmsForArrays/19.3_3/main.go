package main

import (
	"fmt"
	"math"
)

const size int = 10

func twoSmallestNumbersI(input [size]int) (int, int) {
	minIntI := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < minIntI {
			minIntI = input[i]
		}
	}
	minIntII := math.MaxInt
	for i := 0; i < len(input); i++ {
		if input[i] < minIntII && input[i] > minIntI {
			minIntII = input[i]
		}
	}
	return minIntI, minIntII
}

func twoSmallestNumbersII(input [size]int) (int, int) {
	minIntI := math.MaxInt
	minIntII := math.MaxInt
	for i := 0; i < len(input); i++ {
		if input[i] < minIntI {
			minIntII = minIntI
			minIntI = input[i]
		} else if input[i] < minIntII {
			minIntII = input[i]
		}
	}
	return minIntI, minIntII
}

func twoLargestNumbersI(input [size]int) (int, int) {
	maxIntI := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > maxIntI {
			maxIntI = input[i]
		}
	}
	maxIntII := math.MinInt
	for i := 0; i < len(input); i++ {
		if input[i] > maxIntII && input[i] < maxIntI {
			maxIntII = input[i]
		}
	}
	return maxIntI, maxIntII
}

func twoLargestNumbersII(input [size]int) (int, int) {
	maxIntI := math.MinInt
	maxIntII := math.MinInt
	for i := 0; i < len(input); i++ {
		if input[i] > maxIntI {
			maxIntII = maxIntI
			maxIntI = input[i]
		} else if input[i] > maxIntII {
			maxIntII = input[i]
		}
	}
	return maxIntI, maxIntII
}

func main() {
	var arr = [size]int{2, 3, 17, 9, 64, 19, 21, 55, 58, 2}
	fmt.Println(twoSmallestNumbersI(arr))
	fmt.Println(twoSmallestNumbersII(arr))
	fmt.Println(twoLargestNumbersI(arr))
	fmt.Println(twoLargestNumbersII(arr))
}
