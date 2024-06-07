package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		var i int
		for {
			i++
			select {
			case <-exit:
				fmt.Println("exit.")
				return
			default:
				fmt.Println(i * i)
				time.Sleep(time.Second)
			}
		}
	}()
	wg.Wait()
}
