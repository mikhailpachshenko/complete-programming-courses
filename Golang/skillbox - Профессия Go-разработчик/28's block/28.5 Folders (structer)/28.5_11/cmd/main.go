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
		student := s.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scan(&student.Name, &student.Age, &student.Grade)
		if err == io.EOF {
			break
		}
		if err := storage.Put(student); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
	fmt.Println("Хранилище студентов:")
	for _, v := range storage.Get() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
