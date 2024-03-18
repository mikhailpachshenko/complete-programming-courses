package model

import (
	"fmt"
)

func (u *User) ToString() string {
	str := fmt.Sprintf("UserID: %d; Name user: %s; Age: %d; Friends:", u.ID, u.Name, u.Age)
	if len(u.Friends) == 0 {
		str += "\n"
	}
	for i, id := range u.Friends {
		if i == len(u.Friends)-1 {
			str += " UserID: " + fmt.Sprint(id) + ".\n"
		} else if i == 0 {
		} else {
			str += " UserID: " + fmt.Sprint(id) + ","
		}
	}
	return str
}
