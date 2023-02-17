package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringIn := "10 10s 20cd 20s20 20 30 30c c30dd"
	arrOfNum := []int{}
	stringSplit := strings.Split(stringIn, " ")
	for _, val := range stringSplit {
		if digit, err := strconv.Atoi(val); err == nil {
			arrOfNum = append(arrOfNum, digit)
		}
	}
	fmt.Println(arrOfNum)
}
