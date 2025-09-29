package main

import (
	"fmt"
	"log"
	"net/http"

	routers "proxy/pkg"
)

const proxyAddr string = "localhost:9000"

func main() {

	fmt.Println("Proxy service started:", proxyAddr)

	/*  Шаблон запроса: curl -X POST -d '{"Name": "Name Surname", "Age": 33, "Friends": []}' http://localhost:9000/create */
	http.HandleFunc("/create", routers.HandleCreateUser)

	/* Шаблон запроса: curl -X POST -d '{"source_id": 1, "target_id": 2}' http://localhost:9000/make_friends */
	http.HandleFunc("/make_friends", routers.HandleMakeFriends)

	/* Шаблон запроса: curl -DELETE -d '{"target_id": 1}' http://localhost:9000/remove_user */
	http.HandleFunc("/remove_user", routers.HandleRemoveUser)

	/* Шаблон запроса: curl -X GET http://localhost:9000/get_all */
	http.HandleFunc("/get_all", routers.HandleGetAll)

	/* Шаблон запроса: curl -X GET http://localhost:9000/get/1 */
	http.HandleFunc("/get/", routers.HandleGetUser)

	/* Шаблон запроса: curl -X PUT -d '{"id": 1, "age": 29}' http://localhost:9000/new_age */
	http.HandleFunc("/new_age", routers.HandleNewAge)

	log.Fatalln(http.ListenAndServe(proxyAddr, nil))
}
