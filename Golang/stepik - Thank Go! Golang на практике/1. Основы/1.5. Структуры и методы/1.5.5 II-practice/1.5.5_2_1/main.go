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

func minlen(length int) validator {
	return func(s string) bool {
		if utf8.RuneCountInString(s) < length {
			return false
		}
		return true
	}
}

func and(funcs ...validator) validator {
	return func(s string) bool {
		for _, v := range funcs {
			if v(s) == false {
				return false
			}
		}
		return true
	}
}

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
func or(funcs ...validator) validator {
	return func(s string) bool {
		for _, v := range funcs {
			if v(s) == true {
				return true
			}
		}
		return false
	}
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	if p.validator(p.value) {
		return true
	}
	return false
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	var s string
	fmt.Scan(&s)
	// валидатор, который проверяет, что пароль содержит буквы и цифры,
	// либо его длина не менее 10 символов
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
