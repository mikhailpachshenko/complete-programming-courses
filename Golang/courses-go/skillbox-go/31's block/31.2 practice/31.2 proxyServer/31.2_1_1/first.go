package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const addrI string = "localhost:8080"

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(addrI, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	textBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("first:", string(textBytes))
	defer r.Body.Close()

	text := string(textBytes)
	response := "1 instance: " + text + ".\n"

	if _, err = w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
