package main

import (
	"fmt"
	"io"
	"module/pkg/repo/students_repo"
	"module/pkg/student"
)

func main() {
	storage := students_repo.NewStudentsStorage()
	for {
		new := student.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scan(&new.Name, &new.Age, &new.Grade)
		if err == io.EOF {
			break
		}

		if err := storage.Put(new); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
	fmt.Println("Хранилище студентов")
	for _, v := range storage.Get() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
