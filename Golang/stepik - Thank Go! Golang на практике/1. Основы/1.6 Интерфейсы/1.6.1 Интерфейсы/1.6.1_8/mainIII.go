package main

import "fmt"

func aBool(a []int) bool {
	ok := true
	for i, _ := range a {
		ok = i+1 != len(a)
		fmt.Println(ok)
	}
	return ok
}

func val(a []int) int {
	digit := 0
	for _, val := range a {
		digit = val
	}
	return digit
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	digitMax := 0
	for !aBool(a) {
		curr := val(a)
		fmt.Println(curr)
		if curr > digitMax {
			digitMax = curr
		}
	}
	fmt.Println(digitMax)
}
