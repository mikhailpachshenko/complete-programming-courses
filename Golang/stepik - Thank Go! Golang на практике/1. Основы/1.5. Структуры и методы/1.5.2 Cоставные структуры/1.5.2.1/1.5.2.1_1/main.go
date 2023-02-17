package main

import "fmt"

type person struct {
	firstName  string
	secondName string
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
	b := book{
		title: "The Majik Gopher",
		author: person{
			firstName:  "Christopher",
			secondName: "Swanson",
		},
	}
	fmt.Println(b)

	u := user{
		name: "Chris",
		karma: struct {
			value int
			title string
		}{
			value: 100,
			title: "^-^",
		},
	}
	fmt.Printf("%+v\n", u)

	chris := user{
		name: "Chris",
	}
	c := comment{
		text:   "Gophers are awesome!",
		author: &chris,
	}
	fmt.Printf("%+v\n", c)
}
