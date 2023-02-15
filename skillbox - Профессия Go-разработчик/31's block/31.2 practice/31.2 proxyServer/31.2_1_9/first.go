package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr string = "localhost:8080"

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
	response := "1 instance: " + text

	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
