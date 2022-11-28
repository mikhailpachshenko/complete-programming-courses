package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	in := scanner(&wg)
	middle := squaring(in, &wg)
	out := multiplier(middle, &wg)
	receiver(out, &wg)

	wg.Wait()
	fmt.Println("Конец!")
}

func scanner(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Отправитель заканчивает работу.")
			wg.Done()
		}()
		defer func() {
			fmt.Println("Отправитель закрывает канал.")
			close(out)
		}()
		var scan string
		var digit int
		for {
			_, err := fmt.Scan(&scan)
			if err != nil {
				log.Println()
				continue
			}
			digit, err = strconv.Atoi(scan)
			if err != nil {
				if scan == "stop" {
					break
				}
				log.Println(err)
				continue
			}
			fmt.Println("Отправитель принял и отправляет:", digit)
			out <- digit
			time.Sleep(time.Second)
		}
	}()
	return out
}

func squaring(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Возводитель в квадрат заканчивает работу.")
			wg.Done()
		}()
		defer func() {
			fmt.Println("Возводитель в квадрат закрывает канал.")
			close(out)
		}()
		for val := range in {
			res := val * val
			fmt.Println("Возводитель в квадрат отправляет:", res)
			out <- res
		}
	}()
	return out
}

func multiplier(in chan int, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Множитель заканчивает работу.")
			wg.Done()
		}()
		defer func() {
			fmt.Println("Множитель закрывает канал.")
			close(out)
		}()
		for val := range in {
			res := val * 2
			fmt.Println("Множитель отправляет:", res)
			out <- res
		}
	}()
	return out
}

func receiver(in chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("Получатель заканчивает работу.")
			wg.Done()
		}()
		for val := range in {
			fmt.Println("Получатель принял:", val)
		}
	}()
}
