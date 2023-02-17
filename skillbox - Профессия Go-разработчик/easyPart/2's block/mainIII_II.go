package main

import "fmt"

func main() {
	baseRate := 5
	squarePlace := 43
	residents := 4

	result := baseRate * squarePlace * residents

	fmt.Print("Стоимость отопления для площади: ", squarePlace, " и количества жильцов: ", residents, " составит: ", result, "\n")
}
