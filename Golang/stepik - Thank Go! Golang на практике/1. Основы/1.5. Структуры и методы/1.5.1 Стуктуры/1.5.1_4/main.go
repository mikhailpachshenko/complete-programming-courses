package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := person{name: "Mikhail", age: 20}
	fmt.Println(a)
	a.name = "Mike"
	fmt.Println(a)

	var b person
	b.name = "Sasha"
	b.age = 21
	fmt.Println(b)
	b.age = 22
	fmt.Println(b)

	var c person
	fmt.Println(c)

	d := person{"Maksim", 30}
	d.name = "Max"
	fmt.Println(d)

	e := &person{name: "Stanislav", age: 29}
	fmt.Println(e)
	e.name = "Stas"
	fmt.Println((*e).name)
	fmt.Println(e.name)
	fmt.Println(e)

	f := &person{name: "Danil", age: 40}
	fmt.Println(f)

	g := newPerson("Ol'ga")
	fmt.Println(g)

	h := someNewPerson("Nastia")
	fmt.Println(h)

	i := makePerson("Yana")
	fmt.Println(i)
}

func newPerson(name string) *person {
	p := person{
		name: name,
		age:  42,
	}
	return &p
}

func someNewPerson(name string) *person {
	return &person{
		name: name,
		age:  32,
	}
}

func makePerson(name string) person {
	return person{
		name: name,
		age:  22,
	}
}
