package main

import (
	"fmt"
	"unicode"
)

type inn string

// type otherId inn

func (i inn) isValid() bool {
	if len(i) != 12 {
		return false
	}
	for _, char := range i {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	inn1 := inn("121311184921")
	fmt.Println("inn:", inn1, ", is valid:", inn1.isValid())

	inn2 := inn("ohmyinn12345")
	fmt.Println("inn:", inn2, ", is valid:", inn2.isValid())

	/*other := otherId("111201284667")
	fmt.Println("otherId:", other, ", is valid:", other.isValid())
	// other.isValid undefined */
}
