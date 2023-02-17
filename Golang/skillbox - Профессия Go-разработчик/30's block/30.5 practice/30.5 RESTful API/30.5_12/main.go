package main

import (
	"fmt"
	"net/http"
)

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

type MakeFriends struct {
	ID1 string `json:"source_id"`
	ID2 string `json:"target_id"`
}

type NewAge struct {
	Name   string `json:"name"`
	NewAge int    `json:"new age"`
}

func (u User) toString() string {
	s := fmt.Sprintf("UserID is %s, Name is %s, Age is %d, Friends: ", u.ID, u.Name, u.Age)
	for i, v := range u.Friends {
		if i == len(u.Friends)-1 {
			s += fmt.Sprint(v.Name) + ".\n"
		} else if len(u.Friends) == 0 {
			s += "\n"
			break
		} else {
			s += fmt.Sprintf("%s, ", v.Name)
		}
	}
	return s
}

type usersStorage struct {
	database map[string]*User
}

const addr string = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	us := usersStorage{make(map[string]*User)}

	mux.HandleFunc("/create", us.Create)
	mux.HandleFunc("/make_friends", us.makeFriends)
	mux.HandleFunc("/delete_user", us.deleteUser)
	mux.HandleFunc("/get_friends", us.getFriends)
	mux.HandleFunc("/user_id", us.changeAge)

	fmt.Println("server is running")
	http.ListenAndServe(addr, mux)
}
