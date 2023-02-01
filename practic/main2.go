package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Eve"
	b := "ABBA"
	c := "mama"
	d := "Madam, i'm Adam"

	fmt.Println(a, b, c, d)

	fmt.Println(palindrome(a))
	fmt.Println(palindrome(b))
	fmt.Println(palindrome(c))
	fmt.Println(palindrome(d))

}

func palindrome(in string) bool {
	lowerCase := strings.ToLower(in)
	lowerCaseCut := strings.ReplaceAll(lowerCase, " ", "")
	lowerCaseCut = strings.ReplaceAll(lowerCaseCut, ",", "")
	lowerCaseCut = strings.ReplaceAll(lowerCaseCut, "'", "")
	fmt.Println("lowerCaseCut:", lowerCaseCut)
	newString := ""
	for i := len(lowerCaseCut) - 1; i >= 0; i-- {
		newString += string(lowerCaseCut[i])
	}
	fmt.Println("newString:", newString)
	if lowerCaseCut == newString {
		return true
	}
	return false
}
