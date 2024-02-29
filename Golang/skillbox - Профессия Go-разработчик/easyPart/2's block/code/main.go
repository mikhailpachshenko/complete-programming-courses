// Разбор первой команды

package main

/*
Так тоже можно выделять зону под комментарий
*/

import "fmt"

func main() {
	var someVariable int = 42
	fmt.Print("Hello world\n", "the beast\n", 42, "\n") // Функция выводит на экран то что нам нужно

	fmt.Println("Hello World")
	fmt.Println("the beast")
	fmt.Println(42)
	fmt.Print("Значение переменной \"someVariable\": ", someVariable)
}
