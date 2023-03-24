package main

import (
	"fmt"
)

func main() {
	var arrI [4]int
	var arrII = [5]int{1, 2, 3, 4, 5}
	arrIII := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arrI, arrI, arrI)
	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arrII, arrII, arrII)
	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arrIII, arrIII, arrIII)

	for i := 0; i < len(arrIII); i++ {
		fmt.Println(arrIII[i])
	}
}
