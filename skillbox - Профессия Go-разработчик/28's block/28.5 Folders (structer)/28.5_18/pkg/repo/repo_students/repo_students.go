package repo_students

import (
	"fmt"
	s "module/pkg/student"
)

type StudentStorage struct {
	database map[string]*s.Student
}

func NewStudentStorage() *StudentStorage {
	return &StudentStorage{
		database: make(map[string]*s.Student),
	}
}

func (ss *StudentStorage) Put(in *s.Student) error {
	if in == nil {
		return fmt.Errorf("Information about student is empty.")
	}
	if in.Name == "" {
		return fmt.Errorf("Row name about student is empty.")
	}
	ss.database[in.Name] = in
	return nil
}

func (ss StudentStorage) Get() []*s.Student {
	res := []*s.Student{}
	for _, v := range ss.database {
		res = append(res, v)
	}
	return res
}
