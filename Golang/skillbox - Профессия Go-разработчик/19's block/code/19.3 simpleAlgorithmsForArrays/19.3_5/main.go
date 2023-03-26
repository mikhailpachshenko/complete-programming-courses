package main

import (
	"fmt"
	"math"
)

const size int = 10

func twoSmallestNumbersI(in [size]int) (int, int) {
	minIntI := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] < minIntI {
			minIntI = in[i]
		}
	}
	minIntII := math.MaxInt
	for i := 0; i < len(in); i++ {
		if in[i] < minIntII && in[i] > minIntI {
			minIntII = in[i]
		}
	}
	return minIntI, minIntII
}

func twoLargestNumbersI(in [size]int) (int, int) {
	maxIntI := in[0]
	for i := 0; i < len(in); i++ {
		if in[i] > maxIntI {
			maxIntI = in[i]
		}
	}
	maxIntII := math.MinInt
	for i := 0; i < len(in); i++ {
		if in[i] > maxIntII && in[i] < maxIntI {
			maxIntII = in[i]
		}
	}
	return maxIntI, maxIntII
}

func twoSmallestNumbersII(in [size]int) (int, int) {
	minIntI := math.MaxInt
	minIntII := math.MaxInt
	for i := 0; i < len(in); i++ {
		if in[i] < minIntI {
			minIntII = minIntI
			minIntI = in[i]
		} else if in[i] < minIntII {
			minIntII = in[i]
		}
	}
	return minIntI, minIntII
}

func twoLargestNumbersII(in [size]int) (int, int) {
	maxIntI := math.MinInt
	maxIntII := math.MinInt
	for i := 0; i < len(in); i++ {
		if in[i] > maxIntI {
			maxIntII = maxIntI
			maxIntI = in[i]
		} else if in[i] > maxIntII {
			maxIntII = in[i]
		}
	}
	return maxIntI, maxIntII
}

func main() {
	var arr = [size]int{1, 3, 8, 9, 13, 0, 29, 28, 29, 1}
	fmt.Println(twoSmallestNumbersI(arr))
	fmt.Println(twoSmallestNumbersII(arr))
	fmt.Println(twoLargestNumbersI(arr))
	fmt.Println(twoLargestNumbersII(arr))
}
