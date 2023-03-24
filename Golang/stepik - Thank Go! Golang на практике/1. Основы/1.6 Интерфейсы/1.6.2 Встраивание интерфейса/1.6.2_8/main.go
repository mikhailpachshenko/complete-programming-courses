package main

import (
	"fmt"
)

type counter struct {
	val uint
}

func (c *counter) increment() {
	c.val++
}

func (c *counter) value() uint {
	return c.val
}

type Counter interface {
	increment()
	value() uint
}

type usage struct {
	service string
	Counter
}

func newUsage(service string) *usage {
	return &usage{service, &counter{}}
}

func main() {
	u := newUsage("find")
	u.increment()
	u.increment()
	u.increment()
	u.increment()
	fmt.Printf("%s, usage: %d.\n", u.service, u.value())
}
