package students_repo

import (
	"fmt"
	"module/pkg/student"
)

type StudentsStorage struct {
	database map[string]*student.Student
}

func NewStudentsStorage() *StudentsStorage {
	return &StudentsStorage{
		database: make(map[string]*student.Student),
	}
}

func (ss *StudentsStorage) Put(in *student.Student) error {
	if in == nil {
		return fmt.Errorf("Inf-on abt students is empty")
	}

	if in.Name == "" {
		return fmt.Errorf("Name student is empty")
	}

	ss.database[in.Name] = in
	return nil
}

func (ss *StudentsStorage) Get() []*student.Student {
	res := []*student.Student{}
	for _, v := range ss.database {
		res = append(res, v)
	}
	return res
}
