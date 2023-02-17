package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go putBook(&wg)
	go deliverBook(&wg)

	wg.Wait()
	burnBook()
}

func putBook(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Складываю книги!")
}

func deliverBook(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Доставляю книги!")
}

func burnBook() {
	fmt.Println("Сжигаю книги!")
}
