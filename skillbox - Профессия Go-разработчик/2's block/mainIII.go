// fmt.Println(2 + 2) // Вывод 4 (обратим внимание на отсутствие кавычек), "+" — оператор

package main

import "fmt"

func main() {
	var a int = 21
	var b int = 22
	result := a + b

	fmt.Println(result)
	fmt.Println("a + b =", result)
	fmt.Print("a + b = ", a+22, "\n")

	a = 2
	b = 3
	c := 4

	result = a + b*c // правила порядка операций действует как в правилах арифметики
	fmt.Println(result)
	/*
	   Go самостоятельно определит тип данного отталкиваясь от значения операнда
	*/

	/*
		7 - 3
		——————
		10 - 6
	*/

	result = (7 - 1) / (10 - 6)
	fmt.Println(result)
}
