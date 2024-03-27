package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"module/pkg/dao"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandlerCreateUser(t *testing.T) {

	type User struct {
		Name    string   `json:"name"`
		Age     string   `json:"age"`
		Friends []string `json:"friends"`
	}
	// !
	type Users struct {
		Users []User `json:"users"`
	}
	// !
	jsonFile, err := os.Open("/Users/mikhailpachshenko/Desktop/sk/semi-final/service_a/pkg/testing/createuser.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	// !
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	// !
	var users Users
	if json.Unmarshal(byteValue, &users); err != nil {
		log.Fatalln(err)
	}
	// !
	for i := 0; i < len(users.Users); i++ {
		jsonData, err := json.Marshal(users.Users[i])
		if err != nil {
			log.Fatalln(err)
		}

		srv := httptest.NewServer(http.HandlerFunc(dao.CreateUser))
		resp, err := http.Post(srv.URL, "application/text", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Log(err)
			t.Fail()
		}

		if resp.StatusCode != 201 {
			fmt.Println("Ошибка: " + users.Users[i].Name)
			t.Fail()
		}

		defer resp.Body.Close()
	}
}

func TestHandlerListUsers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(dao.ListUsers))
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if resp.StatusCode != 200 {
		t.Fail()
	}
}
