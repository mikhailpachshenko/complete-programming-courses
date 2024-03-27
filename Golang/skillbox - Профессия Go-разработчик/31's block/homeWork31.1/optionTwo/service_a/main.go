package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	dao "module/pkg/dao"
)

func main() {

	fmt.Println("Service start on port 8081")

	flag.Parse()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Post("/create", dao.CreateUser)
	r.Post("/make_friends", dao.MakeFriends)
	r.Delete("/user", dao.DeleteUser)
	r.Get("/friends/{user_id}", dao.GetFriends)
	r.Get("/list", dao.ListUsers)
	r.Put("/{user_id}", dao.ChangeAge)
	http.ListenAndServe(":8081", r)
}
