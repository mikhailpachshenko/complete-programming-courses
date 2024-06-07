package main

import (
	"fmt"
)

func main() {
	fc := putBook()
	sc := deliverBook(fc)
	tc := burnBook(sc)

	fmt.Println(<-tc)
}

func putBook() chan string {
	firstChan := make(chan string)
	go func() {
		firstChan <- "Складываю книги!"
	}()
	return firstChan
}

func deliverBook(firstChan chan string) chan string {
	secondChan := make(chan string)
	fmt.Println(<-firstChan)
	go func() {
		secondChan <- "Достааляю книги!"
	}()
	return secondChan
}

func burnBook(secondChan chan string) chan string {
	thirdChan := make(chan string)
	fmt.Println(<-secondChan)
	go func() {
		thirdChan <- "Сжигаю книги!"
	}()
	return thirdChan
}
