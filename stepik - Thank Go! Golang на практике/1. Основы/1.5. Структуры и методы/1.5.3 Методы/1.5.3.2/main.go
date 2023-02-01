package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

func main() {
	cptr := &circle{radius: 5}
	fmt.Println("circle area:", cptr.area())

	r := circle{radius: 3}

	rptr := &r
	fmt.Println(rptr.area())
	fmt.Println(r.area())

	c := *cptr
	fmt.Println(c.area())
	fmt.Println(cptr.area())
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}
