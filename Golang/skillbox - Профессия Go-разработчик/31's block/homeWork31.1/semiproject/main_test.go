package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"semiproject/pkg/model"
	"testing"
)

func TestHandlerCreateUser(t *testing.T) {
	/* открываем тестовый JSON файл + отложенное закрытие os.Open */
	jsonFile, err := os.Open("/Users/mikhailpachshenko/Desktop/31.1 2/semiproject/pkg/test/addII_user.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	/* получаем байты информации с предыдущего шага */
	content, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	/* инициализируем соответствующий тип для Unmarshal */
	usI := model.NewUsersStorage()
	if err := json.Unmarshal(content, &usI); err != nil {
		log.Fatalln(err)
	}

	/* цикл: итерируемся по Store вытаскиваем User и отправляем в Marshal для последующего пропуска информации через тестовый POST запрос*/
	for i := range usI.Store {
		usII := usI
		jsonData, err := json.Marshal(usII.Store[i])
		if err != nil {
			log.Fatalln(err)
		}

		/* запускаем тестовый сервер с необходимым запросом POST c отложенным закрытием */
		srv := httptest.NewServer(http.HandlerFunc(usI.CreateUser))

		resp, err := http.Post(srv.URL, "application/text", bytes.NewBuffer(jsonData))

		if err != nil {
			t.Log(err)
			t.Fail()
		}

		/* Ошибка в случае если пользователь не создан */
		if resp.StatusCode != 201 {
			fmt.Println("Ошибка: " + usII.Store[i].Name + " не прошел проверку по длине имени или по возрасту.")
			t.Fail()
		} else {
			fmt.Println("Пользователь успешно внесен в базу данных")
			fmt.Println()
		}
		defer resp.Body.Close()
	}
}

func TestHandlerListUsers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(model.NewUsersStorage().GetAll))
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if resp.StatusCode != 200 {
		t.Fail()
	}
}
