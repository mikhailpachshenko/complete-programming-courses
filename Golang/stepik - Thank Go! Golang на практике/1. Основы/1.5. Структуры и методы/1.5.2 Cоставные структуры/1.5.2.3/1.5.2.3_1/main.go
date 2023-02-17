package main

import "fmt"

type User struct {
	name string
	age  int
	next *User
	prev *User
}

func main() {
	user1 := User{"user1", 1, nil, nil}
	user2 := User{"user2", 2, nil, nil}
	user3 := User{"user3", 3, nil, nil}

	user1.next = &user2
	user2.next = &user3
	user3.prev = &user2
	user2.prev = &user1

	user := &user1

	for user != nil {
		fmt.Println(user.name, user.age)
		user = user.next
	}

	user = &user3
	for user != nil {
		fmt.Println(user.name, user.age)
		user = user.prev
	}
}
