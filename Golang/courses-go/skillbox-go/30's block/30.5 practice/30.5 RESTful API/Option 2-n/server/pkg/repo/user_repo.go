package user_repo

import (
	user "server/pkg/user"
	"sync"
)

var mutex sync.Mutex

type Store map[int]*user.User

func InitStore() Store {
	return make(map[int]*user.User)
}

func (us Store) Put(u *user.User, wg *sync.WaitGroup) {
	mutex.Lock()
	us[u.Id] = u
	mutex.Unlock()
	wg.Done()
}

/*
func (us Store) Put(u *user.User) []byte {
	us[u.Id] = u
	return []byte("User created ID = " + strconv.Itoa(u.Id) + " \n")
}
*/

func (us Store) Get() []byte {
	response := ""
	for _, user := range us {
		response += user.ToString()
	}
	return []byte(response)
}

func (us Store) Delete(id int) []byte {
	var userName string

	for _, user := range us {
		if user.Id == id {
			delete(us, id)
			userName = user.Name
		}
	}

	for _, user := range us {
		for i := 0; i < len(user.Friends); i++ {
			if user.Friends[i] == id {
				copy(user.Friends[i:], user.Friends[i+1:])
				user.Friends[len(user.Friends)-1] = 0
				user.Friends = user.Friends[:len(user.Friends)-1]
			}
		}
	}

	return []byte("Пользователь " + userName + " успешно удален \n")
}

func (us Store) IdCounter() int {
	var max int = 0
	for _, user := range us {
		if user.Id >= max {
			max = user.Id + 1
		}
	}
	return max
}

func (us Store) Friend(source_id, target_id int) []byte {

	var source_name string
	var target_name string

	for _, user := range us {
		if user.Id == source_id {
			user.Friends = append(user.Friends, target_id)
			source_name = user.Name
		}
		if user.Id == target_id {
			user.Friends = append(user.Friends, source_id)
			target_name = user.Name
		}
	}
	return []byte(source_name + " и " + target_name + " теперь друзья \n")
}

func (us Store) EditAge(id, age int) []byte {
	var answer []byte
	for _, user := range us {
		if user.Id == id {
			user.Age = age
			answer = []byte("Возраст пользователя " + user.Name + " успешно обновлен \n")
		}
	}
	return []byte(answer)
}

func (us Store) GetAllFriends(id int) []byte {

	response := ""
	for _, user := range us {
		if user.Id == id {
			response += user.ToFriends()
		}
	}

	return []byte(response)
}
