package main

import "fmt"

type user struct {
	name   string
	number int
	next   *user
	prev   *user
}

func main() {
	user1 := user{"user1", 1, nil, nil}
	user2 := user{"user2", 2, nil, nil}
	user3 := user{"user3", 3, nil, nil}

	user1.next = &user2
	user2.next = &user3
	user2.prev = &user1
	user3.prev = &user2

	user := &user1

	for user != nil {
		fmt.Println((*user).name, (*user).number)
		user = user.next
	}

	user = &user3
	for user != nil {
		fmt.Println(user.name, user.number)
		user = user.prev
	}
}
