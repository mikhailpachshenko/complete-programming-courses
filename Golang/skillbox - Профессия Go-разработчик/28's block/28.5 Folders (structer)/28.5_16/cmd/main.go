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
		nStudent := s.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scan(&nStudent.Name, &nStudent.Age, &nStudent.Grade)
		if err == io.EOF {
			break
		}
		if err := storage.Put(nStudent); err != nil {
			fmt.Println("error: ", err.Error())
			return
		}
	}
	fmt.Println("Хранилище студентов:")
	for _, v := range storage.Get() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
