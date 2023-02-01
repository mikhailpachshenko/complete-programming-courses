package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) scale(factor int) {
	r.width *= factor
	r.height *= factor
}

func (r *rect) pScale(factor int) {
	r.width *= factor
	r.height *= factor
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("rect area:", r.area())

	fmt.Println("rect before scaling:", r)
	r.scale(2)
	fmt.Println("rect after scaling:", r)

	fmt.Println("rect before pScaling:", r)
	r.pScale(2)
	fmt.Println("rect after pScaling:", r)
}
