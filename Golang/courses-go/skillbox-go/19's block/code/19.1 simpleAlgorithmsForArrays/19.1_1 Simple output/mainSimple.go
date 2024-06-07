package main

import (
	"fmt"
)

func main() {
	//var numbersI [5]int
	//var numbersII = [5]int{10, 49, 48, 20, 39}
	numbersII := [5]int{10, 49, 48, 20, 39}
	fmt.Println("index 0: number =", numbersII[0], "index 2: number =", numbersII[2])

	for i := 0; i < len(numbersII); i++ {
		fmt.Printf("Index = %d, number = %d.\n", i, numbersII[i])
	}
}
