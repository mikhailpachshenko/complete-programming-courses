package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type validator func(s string) bool

func digits(s string) bool {
	for _, v := range s {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

func letters(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

func minlen(lenght int) validator {
	return func(s string) bool {
		return utf8.RuneCountInString(s) >= lenght
	}
}

func and(funcs ...validator) validator {
	return func(s string) bool {
		for _, v := range funcs {
			if !v(s) {
				return false
			}
		}
		return true
	}
}

func or(funcs ...validator) validator {
	return func(s string) bool {
		for _, v := range funcs {
			if v(s) {
				return true
			}
		}
		return false
	}
}

type password struct {
	value string
	validator
}

func (p *password) IsValid() bool {
	return p.validator(p.value)
}

func main() {
	var s string
	fmt.Scan(&s)
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.IsValid())
}
