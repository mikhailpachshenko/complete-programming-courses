package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringIn := "20 1x3 42 12ca 200 2O1"
	arrOfNum := []int{}
	stringSplit := strings.Split(stringIn, " ")
	for _, v := range stringSplit {
		if digit, err := strconv.Atoi(v); err == nil {
			arrOfNum = append(arrOfNum, digit)
		} else if _, err := strconv.Atoi(v); err != nil {
			var stringSome string
			for _, s := range v {
				if digit, err := strconv.Atoi(string(s)); err == nil {
					stringSome += strconv.Itoa(digit)
				}
			}
			intSome, _ := strconv.Atoi(stringSome)
			arrOfNum = append(arrOfNum, intSome)
		}
	}
	fmt.Println(arrOfNum)
}
