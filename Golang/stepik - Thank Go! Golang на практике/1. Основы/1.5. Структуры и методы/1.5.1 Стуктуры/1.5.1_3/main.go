package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	a := person{"Mike", 20}
	fmt.Println(a)

	b := person{name: "Mikhail", age: 30}
	fmt.Println(b)

	c := person{name: "Sam"}
	fmt.Println(c)

	d := &person{"Ira", 25}
	fmt.Println(d)

	e := person{"Sveta", 26}
	fmt.Println(e)
	e.name = "Sasha"
	fmt.Println(e)

	f := &person{"Maksim", 33}
	fmt.Println((*f).name)
	fmt.Println(f.name)

	g := newPerson("Kirill")
	fmt.Println(g)

	h := nNewPerson("Nikita")
	fmt.Println(h)

	i := makePerson("Arseniy")
	fmt.Println(i)
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func nNewPerson(name string) *person {
	return &person{
		name: name,
		age:  42,
	}
}

func makePerson(name string) person {
	return person{
		name: name,
		age:  30,
	}
}
