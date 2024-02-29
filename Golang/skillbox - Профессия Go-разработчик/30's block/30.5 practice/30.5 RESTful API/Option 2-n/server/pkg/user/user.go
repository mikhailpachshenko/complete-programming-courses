package user

import "fmt"

type User struct {
	Id   int
	Name string
	Age  int

	Friends []int
}

func NewUser(id int, name string, age int) *User {
	return &User{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

func (u *User) ToFriends() string {
	return fmt.Sprintf("list friends - %d \n", u.Friends)
}

func (u *User) ToString() string {
	return fmt.Sprintf("ID = %d User name is %s and age is %d list friends - %d \n", u.Id, u.Name, u.Age, u.Friends)
}
