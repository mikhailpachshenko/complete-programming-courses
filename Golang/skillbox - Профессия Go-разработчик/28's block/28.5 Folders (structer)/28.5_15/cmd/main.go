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
		nS := student.NewStudent()
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scan(&nS.Name, &nS.Age, &nS.Grade)
		if err == io.EOF {
			break
		}
		if err = storage.Put(nS); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
	}
	fmt.Println("Хранилище студнетов:")
	for _, v := range storage.Get() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
