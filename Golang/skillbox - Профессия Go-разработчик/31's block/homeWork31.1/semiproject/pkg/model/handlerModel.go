package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	config "semiproject/pkg/config"

	"github.com/lib/pq"
)

/* Создание нового пользователи */
func (us *UsersStorage) CreateUser(w http.ResponseWriter, r *http.Request) {

	/* Тело запроса */
	if r.Method == "POST" {

		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		/* Проверки при создании пользователя */
		switch {
		case len(u.Name) <= 2:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("A name shorter than two characters"))
			return
		case u.Age <= 0 || u.Age > 120:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Age specified incorrectly, needs to be between 1 and 120"))
			return
		}

		/* Работаем с БД */
		db := config.ConnectionDatabase()
		defer db.Close()

		/* Запросы в БД */
		pgQuery := `INSERT INTO "usersStore" ("name", "age", "friends") VALUES ($1, $2, '{}') RETURNING id`

		var userID int64 // Инициализаия: переменная для хранения, созданного ID базой данных, при получении данных
		if err = db.QueryRow(pgQuery, u.Name, u.Age).Scan(&userID); err != nil {
			log.Fatalln()
		}
		/* Тело ответа о положительном результате */
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User has been created\nID = " + fmt.Sprint(userID) + " " + fmt.Sprint(u.Name) + "\n"))
		log.Print("\nUser has been created\nID = " + fmt.Sprint(userID) + " " + fmt.Sprint(u.Name) + "\n")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

/* Добавляем в списки друзей пользователей друг к другу */
func (us *UsersStorage) MakeFriends(w http.ResponseWriter, r *http.Request) {

	/* Тело запроса */
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var mf MakeFriends
		if err := json.Unmarshal(content, &mf); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		/* Работаем с БД */
		db := config.ConnectionDatabase()
		defer db.Close()

		/* Запросы в БД */
		query_select := `SELECT "id", "name", "age","friends" FROM "usersStore" WHERE id IN ($1,$2)`
		query_update := `UPDATE "usersStore" SET friends = $1 WHERE id = $2`

		/* Вытаскиваем с БД необходимы данные */
		rows, err := db.Query(query_select, mf.ID1, mf.ID2)
		if err != nil {
			log.Fatalln(err)
		}
		var listMF []*User // Инициализация: срез для внесения данных с запроса для последующей работы
		/*  Цикл: в зависимости от количества переменных которые нам необходимо получить, столько раз итерируемся */
		for rows.Next() {
			var u User
			/* получаем данные и вносим в срез(строка ниже) */
			if err = rows.Scan(&u.ID, &u.Name, &u.Age, pq.Array(&u.Friends)); err != nil {
				log.Fatalln(err)
			}
			listMF = append(listMF, &u)
		}

		/* Проверям нет ли пользователей в списках друзей у друг друга */
		if len(listMF) > 1 {
			for _, user := range listMF {
				for _, id := range user.Friends {
					if id == mf.ID1 {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte("Пользователи уже друзья.\n"))
						return
					}
				}
				for _, id := range user.Friends {
					if id == mf.ID2 {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte("Пользователи уже друзья.\n"))
						return
					}
				}
			}

			/* Добавлление в списки друзей, пользователей, друг к другу */
			var userNameI, userNameII string // Инициализация: переменных "Имен" для оповещения о положительном запросе
			for _, userI := range listMF {
				if userI.ID == mf.ID1 {
					userNameI = userI.Name
					for _, userII := range listMF {
						if userII.ID == mf.ID2 {
							userNameII = userII.Name
							/* Вносим в срез друзей первого полльзователя, данные второго пользователя (строка ниже) */
							userI.Friends = append(userI.Friends, userII.ID)
							/* Обновляем срез друзей первого пользователя используя первичный ключ первого пользователя в виде ID (строка ниже) */
							if _, err := db.Exec(query_update, pq.Array(userI.Friends), mf.ID1); err != nil {
								log.Fatalln(err)
							}
							/* Вносим в срез друзей второго полльзователя, данные первого пользователя (строка ниже) */
							userII.Friends = append(userII.Friends, userI.ID)
							/* Обновляем срез друзей второго пользователя используя первичный ключ второго в виде ID (строка ниже) */
							if _, err := db.Exec(query_update, pq.Array(userII.Friends), mf.ID2); err != nil {
								log.Fatalln(err)
							}
						}
					}
				}
			}
			defer rows.Close()

			/* Тело ответа о положительном результате */
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Users has been added to buddy list! " + userNameI + " and " + userNameII + " are now friends.\n"))
			log.Print("\nUsers has been added to buddy list! " + userNameI + " and " + userNameII + " are now friends.\n")
			return

			/* Тело ответа о отрицательном результате */
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

/* Удаление пользователя (строка ниже) */
func (us *UsersStorage) RemoveUser(w http.ResponseWriter, r *http.Request) {

	/* Тело запроса */
	if r.Method == "DELETE" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var rmUser RemoveUser
		if err := json.Unmarshal(content, &rmUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		/* Работаем с БД */
		db := config.ConnectionDatabase()
		defer db.Close()
		/* в rows вносим запрос с базы данных по указанному ID со всеми его колонками */
		rows, err := db.Query(`SELECT "id", "name", "age", "friends" FROM "usersStore" WHERE id = $1`, &rmUser.ID)
		var rmU User // Инициализация: вносим данные при отправленном ID в бд
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				/* вносим данные полученые с SELECT-запроса в переменную  */
				err = rows.Scan(&rmU.ID, &rmU.Name, &rmU.Age, pq.Array(&rmU.Friends))
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		nameRmUser := rmU.Name // переменная для хранения имени удалённого пользователя

		/* Итерируемся по списку друзей пользователя, которого планируем удалить, для того что бы обновить списки друзей у другий пользователей (строка ниже)*/
		for _, userID := range rmU.Friends {
			/* Запрос в БД (ниже строкой): берем ID пользователя что в друзьях у пользователя что планируем удалить (строка ниже) */
			rows, err := db.Query(`SELECT "id", "name", "age", "friends" FROM "usersStore" WHERE id = $1`, userID)
			var testUser User // переменная куда будем вытаскиать данных из запроса в БД
			if err != nil {
				fmt.Println(err)
			} else {
				for rows.Next() {
					/* из запроса в БД заполняем переменную (строка ниже) */
					err = rows.Scan(&testUser.ID, &testUser.Name, &testUser.Age, pq.Array(&testUser.Friends))
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
			newArr := make([]int64, 0) // инициализация: новый срез где будем вносить данные без удаленного пользователя
			/* Цикл: пробегаем по друзьям пользователя где есть пользователь которого планируем удалить (строка ниже) */
			for i, id := range testUser.Friends {
				if id == rmU.ID {
					newArr = append(newArr, testUser.Friends[:i]...)
					newArr = append(newArr, testUser.Friends[i+1:]...)
					testUser.Friends = newArr
					/* Работа с БД: вносим обновлённый срез друзей без удалённого пользователя */
					_, err = db.Exec(`UPDATE "usersStore" SET friends = $1 WHERE id = $2`, pq.Array(newArr), &testUser.ID)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		}
		defer rows.Close()
		_, err = db.Exec(`DELETE FROM "usersStore" WHERE id = $1`, rmU.ID)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User: " + nameRmUser + ", has been removed!\n"))
		log.Print("User: " + nameRmUser + ", has been removed!\n")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		/* Тело запроса (строка ниже) */
		db := config.ConnectionDatabase()
		defer db.Close()

		/* Текст запроса указанных данных по ID  */
		query := `SELECT "id","name","age", "friends" FROM "usersStore" ORDER by id`
		rows, err := db.Query(query)
		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()

		var list []*User // Инициализируемся: срез для внесения полученных данных из БД
		/* Цикл: итерируемся по всей БД */
		for rows.Next() {
			var u User
			err = rows.Scan(&u.ID, &u.Name, &u.Age, pq.Array(&u.Friends))
			if err != nil {
				fmt.Println(err)
			}
			list = append(list, &u)
		}

		/* Проверка и последующая подготовка данных для вывода ответа */
		if len(list) == 0 {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("База данных пуста"))
			return
		} else {
			var response string // строка ответа
			for _, user := range list {
				response += user.ToString()
			}

			/* тело ответа (строка ниже) */
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		str := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}

		/* Тело запроса */
		db := config.ConnectionDatabase()
		defer db.Close()

		/* Текст запроса */
		rows, err := db.Query(`SELECT "id","name","age","friends" FROM "usersStore" WHERE id = $1`, id)
		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()

		var response string
		for rows.Next() {
			var u User
			err = rows.Scan(&u.ID, &u.Name, &u.Age, pq.Array(&u.Friends))
			if err != nil {
				fmt.Println("Такого пользователя не существует.")
			}
			response += u.ToString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (us *UsersStorage) NewAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer r.Body.Close()

		var na NewAge
		if err := json.Unmarshal(content, &na); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(na)

		query := `update "usersStore" set age = $1 where id = $2`
		db := config.ConnectionDatabase()
		defer db.Close()

		_, err = db.Exec(query, na.Age, na.ID)
		if err != nil {
			log.Fatalln(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Users age has been updated.\n"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
