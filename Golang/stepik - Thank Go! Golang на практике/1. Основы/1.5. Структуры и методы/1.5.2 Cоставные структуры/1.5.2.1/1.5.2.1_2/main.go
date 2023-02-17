package main

import "fmt"

type person struct {
	name string
	age  int
}

type book struct {
	title  string
	author person
}

type user struct {
	name  string
	karma struct {
		value int
		title string
	}
}

type comment struct {
	text   string
	author *user
}

func main() {
	a := person{"Mikhail", 30}
	aBook := book{"Good things", a}
	fmt.Println(aBook)
	fmt.Println(a)

	bBook := book{
		title: "Stay back",
		author: person{
			name: "Sam",
			age:  29,
		},
	}
	fmt.Println(bBook)

	c := struct {
		value int
		title string
	}{
		value: 20,
		title: "Some new text",
	}

	cUser := user{"Sasha", c}
	fmt.Println(cUser)

	dUser := user{
		name: "Dima",
		karma: struct {
			value int
			title string
		}{
			value: 100,
			title: "before auto",
		},
	}
	fmt.Printf("%+v\n", dUser)

	e := user{
		name: "Chris",
	}

	eComment := comment{
		text:   "new text",
		author: &e,
	}
	fmt.Printf("%+v\n", eComment)
}
