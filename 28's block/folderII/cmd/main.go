package main

import (
	"fmt"
	"io"
	ss "module/pkg/repo/students_repo"
	s "module/pkg/student"
)

func main() {
	storage := ss.NewStudentsStorage()
	for {
		new := s.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента.")
		_, err := fmt.Scan(&new.Name, &new.Age, &new.Grade)
		if err == io.EOF {
			break
		}

		if err := storage.Put(new); err != nil {
			fmt.Println("error", err.Error())
			return
		}
	}
	fmt.Println("Студенты из хранилища.")
	for _, in := range storage.Get() {
		fmt.Println(in.Name, in.Age, in.Grade)
	}
}
