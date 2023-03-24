package main

import (
	"fmt"
)

func main() {
	var arrI [5]int
	var arrII = [5]int{1, 2, 3, 4, 5}
	arrIII := [5]int{2, 3, 4, 5, 6}

	fmt.Printf("%d, %t, %T, %+#v\n", arrI, arrI, arrI, arrI)
	fmt.Printf("%d, %t, %T, %#v\n", arrII, arrII, arrII, arrII)
	fmt.Printf("%d, %t, %T, %#+v\n", arrIII, arrIII, arrIII, arrIII)
	fmt.Printf("index 0 of arrII: %d, index 3 of arrII: %d\n", arrII[0], arrII[3])

	for i := 0; i < len(arrIII); i++ {
		fmt.Println(arrIII[i])
	}
}
