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
		fmt.Println("Введите имя, возраст и оценка студентов:")
		_, err := fmt.Scan(&new.Name, &new.Age, &new.Grade)
		if err == io.EOF {
			break
		}

		if err := storage.Put(new); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
	fmt.Println("Студенты из хранилища:")
	for _, in := range storage.Get() {
		fmt.Println(in.Name, in.Age, in.Grade)
	}
}
