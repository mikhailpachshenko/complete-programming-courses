package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
