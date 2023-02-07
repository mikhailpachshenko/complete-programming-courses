package repo_students

import (
	"fmt"
	s "module/pkg/student"
)

type StudentsStorage struct {
	database map[string]*s.Student
}

func NewStudentsStorage() *StudentsStorage {
	return &StudentsStorage{
		database: make(map[string]*s.Student),
	}
}

func (ss *StudentsStorage) Put(in *s.Student) error {
	if in == nil {
		return fmt.Errorf("Information about student is empty.")
	}
	if in.Name == "" {
		return fmt.Errorf("Row name to student is empty.")
	}
	ss.database[in.Name] = in
	return nil
}

func (ss StudentsStorage) Get() []*s.Student {
	res := []*s.Student{}
	for _, v := range ss.database {
		res = append(res, v)
	}
	return res
}
