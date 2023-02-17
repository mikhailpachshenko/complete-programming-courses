package main

import (
	"fmt"
	"io"
	ss "module/pkg/repo/repo_students"
	s "module/pkg/student"
)

func main() {
	storage := ss.NewStudentStorage()
	for {
		newStudent := s.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scan(&newStudent.Name, &newStudent.Age, &newStudent.Grade)
		if err == io.EOF {
			break
		}
		if err = storage.Put(newStudent); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
	fmt.Println("Хранилище студентов:")
	for _, v := range storage.Get() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
