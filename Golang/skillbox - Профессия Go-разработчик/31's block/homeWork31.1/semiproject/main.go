package main

import (
	"fmt"
	"log"
	"net/http"

	"semiproject/pkg/model"
)

const addr string = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	us := model.NewUsersStorage()

	/*  Шаблон запроса: curl -X POST -d '{"Name": "Name Surname", "Age": 33, "Friends": []}' http://localhost:8080/create */
	mux.HandleFunc("/create", us.CreateUser)

	/* Шаблон запроса: curl -X POST -d '{"source_id": 1, "target_id": 2}' http://localhost:8080/make_friends */
	mux.HandleFunc("/make_friends", us.MakeFriends)

	/* Шаблон запроса: curl -X DELETE -d '{"target_id": 1}' http://localhost:8080/remove_user */
	mux.HandleFunc("/remove_user", us.RemoveUser)

	/* Шаблон запроса: curl -X GET http://localhost:8080/get_all */
	mux.HandleFunc("/get_all", us.GetAll) // Получаем весь список пользователей из БД

	/* Шаблон запроса: curl -X GET: http://localhost:8080/get/1 */
	mux.HandleFunc("/get/", us.GetUser) // Получаем данные одного пользователя

	/* Шаблон запроса: curl -X PUT -d '{"id": 1, "age": 29}' http://localhost:8080/new_age */
	mux.HandleFunc("/new_age", us.NewAge)

	fmt.Println("server is running 8080")
	log.Fatalln(http.ListenAndServe(addr, mux))
}
