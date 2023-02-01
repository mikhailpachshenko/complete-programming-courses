package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	bob := person{"Bob", 20}
	fmt.Println(bob) // {Bob 20}

	alice := person{name: "Alice", age: 30}
	fmt.Println(alice) // {Alice 30}

	fred := person{name: "Fred"}
	fmt.Println(fred) // {Fred 0}

	annptr := &person{name: "Ann", age: 40}
	fmt.Println(annptr) // &{Ann 40}

	john := newPerson("John")
	fmt.Println(john) // &{John 42}

	sean := person{name: "Sean", age: 50}
	fmt.Println(sean.name) // Sean

	sven := &person{name: "Sven", age: 50}
	fmt.Println((*sven).age) // 50
	fmt.Println(sven.age)    // 50

	sven.name = "sven"
	sven.age = 51
	fmt.Println(sven)
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func makePerson(name string) person {
	p := person{name: name}
	p.age = 42
	return p
}
