package main

import (
	"fmt"
)

func main() {
	dataChan := getData()
	for val := range dataChan {
		fmt.Println(val)
	}
}

func getData() <-chan int {
	dataChan := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()
	return dataChan
}
