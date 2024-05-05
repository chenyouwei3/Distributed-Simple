package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	Id     int
	Name   string
	Grades []Grade
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

var (
	studentsMutex sync.Mutex
	students      Students
)

func (ss Students) GetByID(id int) (*Student, error) {
	for i := range ss {
		if ss[i].Id == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("找不到学生id:%d", id)
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
