package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr string = "localhost:8001"

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(textBytes)
	response := "2 instance: " + text + ".\n"

	if _, err = w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
