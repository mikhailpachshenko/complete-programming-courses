package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"
	var printNumber []int
	splitInS := strings.Split(s, " ")
	for _, v := range splitInS {
		if digit, err := strconv.Atoi(v); err == nil {
			printNumber = append(printNumber, digit)
		}
	}
	fmt.Println(printNumber)
}
