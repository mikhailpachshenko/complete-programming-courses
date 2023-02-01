package main

import "fmt"

func main() {
	in := "abcdefg"
	out := ""

	for i := len(in) - 1; i >= 0; i-- {
		out += string(in[i])
	}

	fmt.Println(out)
}
