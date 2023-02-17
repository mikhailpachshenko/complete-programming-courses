package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringIn := "a10 10 20b 20 30c30 30 dd"
	arrOfNum := []int{}
	stringSplit := strings.Split(stringIn, " ")
	for _, v := range stringSplit {
		if digit, err := strconv.Atoi(v); err == nil {
			arrOfNum = append(arrOfNum, digit)
		}
	}
	fmt.Println(arrOfNum)
}
