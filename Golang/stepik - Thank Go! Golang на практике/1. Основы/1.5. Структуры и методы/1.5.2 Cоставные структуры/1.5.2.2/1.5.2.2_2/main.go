package main

import "fmt"

type person struct {
	firstName  string
	secondName string
}

type book struct {
	title  string
	author person
}

type pBook struct {
	title  string
	author *person
}

func main() {
	a := person{"Mikhail", "Pachshenko"}
	fmt.Println("\"a\" before use funcBook:", a)
	aBook := book{"Some name", a}
	aBook.author.firstName = "Igor"
	fmt.Println("\"a\" after use funcBook:", a)
	fmt.Println("aBook change firstName:", aBook)

	b := person{"Kirill", "Khivincev"}
	fmt.Println("\"b\" before use pBook:", b)
	bBook := pBook{"some new title to func pBook", &b}
	bBook.author.firstName = "Nikita"
	fmt.Println("\"b\" after use pBook:", b)
	fmt.Println("bBook change firstName:", bBook.author.firstName)
}
