package main

import (
	"fmt"
	"log"
	"net/http"

	routers "proxy/pkg"
)

const proxyAddr string = "localhost:9000"

func main() {

	fmt.Println("Proxy service started: " + proxyAddr)

	http.HandleFunc("/create", routers.HandleCreateUser)
	http.HandleFunc("/make_friends", routers.HandleMakeFriends)
	http.HandleFunc("/user", routers.HandleDeleteUser)
	http.HandleFunc("/friends/", routers.HandleGetFriends)
	http.HandleFunc("/list", routers.HandleList)
	http.HandleFunc("/", routers.HandleChangeAge)

	log.Fatalln(http.ListenAndServe(proxyAddr, nil))

}
