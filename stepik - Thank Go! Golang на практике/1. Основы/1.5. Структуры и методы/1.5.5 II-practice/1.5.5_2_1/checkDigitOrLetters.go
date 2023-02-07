package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func digitsRegexp(s string) bool {
	return regexp.MustCompile(`\d`).MatchString(s)
}

func digitsI(s string) bool {
	for _, v := range s {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println("didgitsRegext:", digitsRegexp(s))
	fmt.Println("digitsI:", digitsI(s))
}
