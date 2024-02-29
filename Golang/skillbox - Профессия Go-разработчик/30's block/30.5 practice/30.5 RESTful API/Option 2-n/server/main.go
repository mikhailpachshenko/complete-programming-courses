package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	user_repo "server/pkg/repo"
	user "server/pkg/user"
)

type UserParrent struct {
	Id      int
	Name    string
	Age     string
	Friends []int
}

type Friendship struct {
	Source_id string `json:"source_id"`
	Target_id string `json:"target_id"`
}

type RemoveUser struct {
	Target_id string `json:"target_id"`
}

type Age struct {
	Newage string `json:"new age"`
}

type service struct {
	user_repo.Store
}

func main() {

	flag.Parse()
	r := chi.NewRouter()
	us := user_repo.InitStore()
	srv := service{us}

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Post("/create", srv.CreateUser)
	r.Post("/make_friends", srv.MakeFriends)
	r.Delete("/user", srv.DeleteUser)
	r.Get("/friends/{user_id}", srv.GetFriends)
	r.Get("/get", srv.GetAll)
	r.Put("/{user_id}", srv.ChangeAge)

	http.ListenAndServe(":3333", r)
}

// Задание №1. Обработчик создания пользователя.
func (s *service) CreateUser(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	content, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var u UserParrent
	if err := json.Unmarshal(content, &u); err != nil {
		fmt.Println(err)
		return
	}

	age, _ := strconv.Atoi(u.Age)
	wg.Add(1)
	id := s.Store.IdCounter()
	go s.Store.Put(user.NewUser(s.Store.IdCounter(), u.Name, age), &wg)
	wg.Wait()
	//w.Write(s.Store.Put(user.NewUser(s.Store.IdCounter(), u.Name, age)))
	w.Write([]byte("User created ID = " + strconv.Itoa(id) + "\n"))
	w.WriteHeader(http.StatusCreated)
}

// Задание №2. Делаем пользователей друзьями.
func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {

	content, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var f Friendship
	if err := json.Unmarshal(content, &f); err != nil {
		fmt.Println(err)
		return
	}

	source_id, _ := strconv.Atoi(f.Source_id)
	target_id, _ := strconv.Atoi(f.Target_id)
	w.Write(s.Store.Friend(source_id, target_id))
	w.WriteHeader(http.StatusOK)
}

// Задание №3. Удаление пользователя
func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {

	content, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var rm RemoveUser
	if err := json.Unmarshal(content, &rm); err != nil {
		fmt.Println(err)
		return
	}

	id, _ := strconv.Atoi(rm.Target_id)
	w.Write(s.Store.Delete(id))
	w.WriteHeader(http.StatusOK)
}

// Задание №4. Возвращает список всех друзей.
func (s *service) GetFriends(w http.ResponseWriter, r *http.Request) {
	var userid int

	content := chi.URLParam(r, "user_id")
	userid, _ = strconv.Atoi(content)

	w.Write([]byte(s.Store.GetAllFriends(userid)))
}

// Дополнительно - вывод всех пользователей.
func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write(s.Store.Get())
}

func (s *service) ChangeAge(w http.ResponseWriter, r *http.Request) {
	// Теперь понятна суть `json: "new age"` привязка со структурой для парса
	userid, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	content, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var na Age
	if err := json.Unmarshal(content, &na); err != nil {
		fmt.Println(err)
		return
	}

	userNewAge, _ := strconv.Atoi(na.Newage)
	w.Write(s.Store.EditAge(userid, userNewAge))
	w.WriteHeader(http.StatusOK)
}
