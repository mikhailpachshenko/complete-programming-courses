package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

const addrProxy string = "localhost:9000"

var (
	counter            int    = 0
	firstInstanceHost  string = "http://localhost:8080"
	secondInstanceHost string = "http://localhost:8081"
)

func main() {
	http.HandleFunc("/", handlerProxy)
	log.Fatalln(http.ListenAndServe(addrProxy, nil))
}

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	textBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	fmt.Println("I:", textBytes, "string:", string(textBytes))
	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(firstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}
		counter++
		fmt.Println("II:", resp)

		textBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println("III:", string(textBytes))
		return
	}
	resp, err := http.Post(secondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}
	counter--
	fmt.Println("IIII:", resp)

	textBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println("V:", string(textBytes))
}
