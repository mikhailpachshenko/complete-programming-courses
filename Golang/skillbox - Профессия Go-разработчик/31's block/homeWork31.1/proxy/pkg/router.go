package routers

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

var (
	counter int = 0
	host    string
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	host, counter = balancer(counter)

	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	resp, err := http.Post(host+r.URL.Path, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusCreated)
	w.Write(textBytes)
}

func HandleMakeFriends(w http.ResponseWriter, r *http.Request) {

	host, counter = balancer(counter)

	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	resp, err := http.Post(host+r.URL.Path, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	w.Write(textBytes)

}

func HandleRemoveUser(w http.ResponseWriter, r *http.Request) {

	host, counter = balancer(counter)

	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	do, err := http.NewRequest(http.MethodDelete, host+r.URL.Path, bytes.NewBuffer([]byte(string(req))))
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	resp, err := client.Do(do)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "text/plain")
	w.Write(textBytes)

}

func HandleGetAll(w http.ResponseWriter, r *http.Request) {

	host, counter = balancer(counter)

	resp, err := http.Get(host + r.URL.Path)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(textBytes)

}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {

	host, counter = balancer(counter)

	resp, err := http.Get(host + r.URL.Path)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write(textBytes)
}

func HandleNewAge(w http.ResponseWriter, r *http.Request) {

	host, counter = balancer(counter)

	req, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	do, err := http.NewRequest(http.MethodPut, host+r.URL.Path, bytes.NewBuffer([]byte(string(req))))
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	resp, err := client.Do(do)
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write(textBytes)
}

func balancer(n int) (hosts string, counter int) {

	listAdrrApplication := [2]string{"http://localhost:8081", "http://localhost:8082"}

	if n == 0 {
		counter = 1
	} else {
		counter = 0
	}

	return listAdrrApplication[n], counter
}
