package main

import (
	"fmt"
)

const size int = 5

func main() {
	var arrI [size]int
	var arrII = [size]int{1, 2, 3, 4, 5}
	arrIII := [size]int{2, 3, 4, 5, 6}

	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arrI, arrI, arrI)
	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arrII, arrII, arrII)
	fmt.Printf("Simple output: %d, Type output: %T, Output like input: %#v\n", arrIII, arrIII, arrIII)
}
