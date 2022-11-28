package students_repo

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
		return fmt.Errorf("Information about student is nil")
	}

	if in.Name == "" {
		return fmt.Errorf("Row name is empty.")
	}

	ss.database[in.Name] = in
	return nil
}

func (ss StudentsStorage) Get() []*s.Student {
	res := []*s.Student{}
	for _, in := range ss.database {
		res = append(res, in)
	}
	return res
}
