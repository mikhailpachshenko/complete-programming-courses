package dao

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lib/pq"

	"github.com/go-chi/chi/v5"

	connection "module/pkg/config"
)

type User struct {
	Id      int
	Name    string
	Age     string
	Friends []string
}

type Friendship struct {
	Source_id string `json:"source_id"`
	Target_id string `json:"target_id"`
}

type RemoveUser struct {
	Target_id string `json:"target_id"`
}

type Age struct {
	Newage string `json:"new age"`
}

// Создание нового пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {

	query := `INSERT INTO "users" ("name", "age", "friends") VALUES ($1, $2, '{}') RETURNING id`

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var u User

		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var userid int

			// Проверка, что имя больше двух символов
			if len(u.Name) < 2 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			age, err := strconv.Atoi(u.Age)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if age < 0 || age > 120 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			db := connection.ConnectionDatabase()
			defer db.Close()

			err = db.QueryRow(query, u.Name, u.Age).Scan(&userid)
			if err != nil {
				log.Fatalln(err)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("User create ID = " + strconv.Itoa(userid) + "\n"))
			}
		}
	}
	defer r.Body.Close()
}

// Делаем пользователей друзьями
func MakeFriends(w http.ResponseWriter, r *http.Request) {

	query_select := `SELECT "id", "name", "age","friends" FROM "users" WHERE id IN ($1,$2)`
	query_update := `UPDATE users SET friends = $1 WHERE id = $2`

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {

		var f Friendship

		if err := json.Unmarshal(content, &f); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {

			var id_user [2]int

			id_user[0], err = strconv.Atoi(f.Source_id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			id_user[1], err = strconv.Atoi(f.Target_id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			db := connection.ConnectionDatabase()
			defer db.Close()

			rows, err := db.Query(query_select, id_user[0], id_user[1])
			if err != nil {
				log.Fatalln(err)
			} else {

				var list []*User

				for rows.Next() {

					var u User
					if err = rows.Scan(&u.Id, &u.Name, &u.Age, pq.Array(&u.Friends)); err != nil {
						log.Fatalln(err)
					}

					list = append(list, &u)
				}

				if len(list) > 1 {
					// Проверка. Являются пользователи друзьями или нет

					for _, s := range list {

						if Find(s.Friends, f.Source_id) {
							w.WriteHeader(http.StatusBadRequest)
							w.Write([]byte("Пользователи уже друзья"))
							return
						}

						if Find(s.Friends, f.Target_id) {
							w.WriteHeader(http.StatusBadRequest)
							w.Write([]byte("Пользователь уже друзья"))
							return
						}
					}

					var username_1, username_2 string
					// Добавляем пользователей в друзья
					for _, s := range list {
						if s.Id != id_user[0] {
							s.Friends = append(s.Friends, f.Source_id)
							if _, e := db.Exec(query_update, pq.Array(s.Friends), f.Target_id); err != nil {
								log.Fatalln(e)
							}
							username_1 = s.Name
						}

						if s.Id != id_user[1] {
							s.Friends = append(s.Friends, f.Target_id)
							if _, e := db.Exec(query_update, pq.Array(s.Friends), f.Source_id); err != nil {
								log.Fatalln(e)
							}
							username_2 = s.Name
						}
					}

					w.WriteHeader(http.StatusOK)
					w.Write([]byte(username_1 + " и " + username_2 + " теперь друзья" + "\n"))

				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			}
			defer rows.Close()
		}
	}

	defer r.Body.Close()
}

// Удаление пользователя по id
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	content, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var ui User
	var id_user RemoveUser

	if err := json.Unmarshal(content, &id_user); err != nil {
		fmt.Println(err)
		return
	} else {

		db := connection.ConnectionDatabase()
		defer db.Close()

		rows, err := db.Query(`SELECT "id", "name", "age", "friends" FROM "users" WHERE id = $1`, &id_user.Target_id)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			for rows.Next() {
				err = rows.Scan(&ui.Id, &ui.Name, &ui.Age, pq.Array(&ui.Friends))
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			for i := range ui.Friends {
				rows, err = db.Query(`SELECT "friends" FROM "users" WHERE id = $1`, &ui.Friends[i])
				if err != nil {
					log.Fatalln(err)
				}
				for rows.Next() {
					var currentListFriends []string
					var newListFriends []string

					if err = rows.Scan(pq.Array(&currentListFriends)); err != nil {
						log.Fatalln(err)
					}
					for j := range currentListFriends {
						if currentListFriends[j] != id_user.Target_id {
							newListFriends = append(newListFriends, currentListFriends[j])
						}
					}

					_, err = db.Exec(`UPDATE users SET friends = $1 WHERE id = $2`, pq.Array(newListFriends), &ui.Friends[i])
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		}

		defer rows.Close()

		s, _ := strconv.Atoi(id_user.Target_id)
		_, err = db.Exec(`DELETE FROM users WHERE id = $1`, s)
		if err != nil {
			log.Fatalln(err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ui.Name))
		}
	}
}

// Показать всех пользователей
func ListUsers(w http.ResponseWriter, r *http.Request) {

	var list []*User
	query := `SELECT "id", "name", "age", "friends" FROM "users" ORDER BY id`

	db := connection.ConnectionDatabase()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Age, pq.Array(&u.Friends))
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, &u)
	}

	if len(list) == 0 {
		w.Write([]byte("Список пустой"))
		w.WriteHeader(http.StatusNoContent)
	} else {
		jsonString, err := json.Marshal(list)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonString)
		}
	}
}

// Вывести список всех друзей пользователя
func GetFriends(w http.ResponseWriter, r *http.Request) {

	var list []*User
	query := `SELECT * FROM users WHERE id IN (SELECT CAST (UNNEST(ARRAY [friends]) AS integer) AS l FROM users WHERE id = $1)`

	content := chi.URLParam(r, "user_id")
	id_user, err := strconv.Atoi(content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {

		db := connection.ConnectionDatabase()
		defer db.Close()

		rows, err := db.Query(query, id_user)
		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()

		for rows.Next() {
			var u User
			if err = rows.Scan(&u.Id, &u.Name, &u.Age, pq.Array(&u.Friends)); err != nil {
				log.Fatalln(err)
			}
			list = append(list, &u)
		}

		jsonString, err := json.Marshal(list)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonString)
	}
}

// Изменить возраст пользователя
func ChangeAge(w http.ResponseWriter, r *http.Request) {

	query := `update users set age = $1 where id = (select id from users u where id = $2)`

	userid, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {

		content, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {

			var age Age
			if err := json.Unmarshal(content, &age); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			userNewAge, err := strconv.Atoi(age.Newage)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			db := connection.ConnectionDatabase()
			defer db.Close()

			_, err = db.Exec(query, userNewAge, userid)
			if err != nil {
				log.Fatalln(err)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Возраст пользователя успешно обновлен" + "\n"))
		}
	}
}

// Функция перебора массива для поиска друзей
func Find(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
