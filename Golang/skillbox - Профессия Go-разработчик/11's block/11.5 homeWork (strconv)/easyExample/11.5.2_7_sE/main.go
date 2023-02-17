package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringIn := "20s 20 31s 200 4x4 21 sd20"
	arrOfNum := []int{}
	stringSplit := strings.Split(stringIn, " ")
	for _, v := range stringSplit {
		if digit, err := strconv.Atoi(v); err == nil {
			arrOfNum = append(arrOfNum, digit)
		}
	}
	fmt.Println(arrOfNum)
}
