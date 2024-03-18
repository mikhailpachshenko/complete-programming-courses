package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const addrII string = "localhost:8081"

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(addrII, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	textBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("second:", string(textBytes))
	defer r.Body.Close()

	text := string(textBytes)
	response := "2 instance: " + text + ".\n"

	if _, err = w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
