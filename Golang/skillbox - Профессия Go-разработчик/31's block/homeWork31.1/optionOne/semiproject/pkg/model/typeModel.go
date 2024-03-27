package model

type User struct {
	Name    string  `json:"name"`
	Age     int64   `json:"age"`
	Friends []int64 `json:"friends"`
	ID      int64   `json:"id"`
}

type UsersStorage struct {
	Store map[string]*User
}

func NewUsersStorage() *UsersStorage {
	return &UsersStorage{
		Store: make(map[string]*User),
	}

}

type MakeFriends struct {
	ID1 int64 `json:"source_id"`
	ID2 int64 `json:"target_id"`
}

type RemoveUser struct {
	ID int64 `json:"target_id"`
}

type NewAge struct {
	ID  int64 `json:"id"`
	Age int64 `json:"age"`
}
