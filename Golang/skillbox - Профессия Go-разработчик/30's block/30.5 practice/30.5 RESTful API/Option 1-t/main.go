package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

type UsersStorage struct {
	Store map[string]*User
}

type MakeFriends struct {
	ID1 string `json:"source_id"`
	ID2 string `json:"target_id"`
}

type RemoveUser struct {
	ID string `json:"target_id"`
}

type ChangeAge struct {
	ID     string `json:"target_id"`
	NewAge int    `json:"new_age"`
}

func (u *User) toString() string {
	str := fmt.Sprintf("UserID: %s; Name user: %s; Age: %d; Friends: ", u.ID, u.Name, u.Age)
	if len(u.Friends) == 0 {
		str += "\n"
	}
	for i, user := range u.Friends {
		if i == len(u.Friends)-1 {
			str += fmt.Sprint(user.Name) + ".\n"
		} else {
			str += fmt.Sprint(user.Name) + ", "
		}
	}
	return str
}

const addr string = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	us := UsersStorage{make(map[string]*User)}

	// Шаблон запроса: curl -X POST -d '{"Name": "Name Surname", "Age": 33, "Friends": []}' http://localhost:8080/create
	mux.HandleFunc("/create", us.creationUser)

	// Шаблон запроса: curl -X POST -d '{"source_id": "c1e834a5", "target_id": "e45c7eb2"}' http://localhost:8080/make_friends
	mux.HandleFunc("/make_friends", us.makeFriends)

	// Шаблон запроса: curl -DELETE -d '{"target_id": "e45c7eb2"}' http://localhost:8080/remove_user
	mux.HandleFunc("/remove_user", us.removeUser)

	// Шаблон запроса в адресной строке браузера: http://localhost:8080/get_all
	// Получаем весь список пользователей
	mux.HandleFunc("/get_all", us.getAll)

	// Шаблон запроса в адресной строке браузера: http://localhost:8080/get/c1e834a5
	// Получаем данные одного пользователя
	mux.HandleFunc("/get/", us.getUserID)

	// Шаблон запроса: curl -X PUT -d '{"target_id": "e45c7eb2", "new_age": 29}' http://localhost:8080/new_age
	mux.HandleFunc("/new_age", us.changeAge)

	fmt.Println("The serveh has been running.")
	http.ListenAndServe(addr, mux)
}

func (us *UsersStorage) creationUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		userID := uuid.New().String()
		splittedID := strings.Split(userID, "-")
		u.ID = splittedID[0]
		us.Store[u.Name] = &u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User has been created:\nName: " + u.Name + "\nID: " + u.ID + "\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) makeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var buddy MakeFriends
		if err := json.Unmarshal(content, &buddy); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var userNameI string
		var userNameII string
		for _, userI := range us.Store {
			if userI.ID == buddy.ID1 {
				userNameI = userI.Name
				for _, userII := range us.Store {
					if userII.ID == buddy.ID2 {
						userNameII = userII.Name
						userI.Friends = append(userI.Friends, userII)
						userII.Friends = append(userII.Friends, userI)
					}
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Users has been added to buddy list! " + userNameI + ", " + userNameII + " are now friends.\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) removeUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var rmvUser RemoveUser
		if err := json.Unmarshal(content, &rmvUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		var nameRmvU string
		out := UsersStorage{make(map[string]*User)}
		func(us *UsersStorage) UsersStorage {
			for _, user := range us.Store {
				arr := make([]*User, 0)
				for i, val := range user.Friends {
					if val.ID == rmvUser.ID {
						nameRmvU = val.Name
						arr = append(arr, user.Friends[:i]...)
						arr = append(arr, user.Friends[i+1:]...)
						user.Friends = arr
					}
				}
				if user.ID != rmvUser.ID {
					out.Store[user.Name] = user
				}
			}
			return out
		}(us)
		us.Store = out.Store
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User: " + nameRmvU + ", has been removed!\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var response string
		for _, user := range us.Store {
			response += user.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) getUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var response string
		id := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		for _, user := range us.Store {
			if id == user.ID {
				response += user.toString()
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) changeAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var cA ChangeAge
		if err := json.Unmarshal(content, &cA); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		var userName string
		for _, user := range us.Store {
			if user.ID == cA.ID {
				userName = user.Name
				user.Age = cA.NewAge
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Возраст " + userName + " успешно изменен!\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
