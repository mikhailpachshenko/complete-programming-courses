package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const addrProxy string = "localhost:9000"

var (
	counter            int    = 0
	firstInstanceHost  string = "http://localhost:8000"
	secondInstanceHost string = "http://localhost:8001"
)

func main() {
	http.HandleFunc("/", proxyHandler)
	log.Fatalln(http.ListenAndServe(addrProxy, nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(firstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}
		counter++

		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(textBytes))
		return
	}

	resp, err := http.Post(secondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}
	counter--

	textBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(textBytes))
}
