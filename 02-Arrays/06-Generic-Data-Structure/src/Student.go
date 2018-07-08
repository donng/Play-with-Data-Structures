package main

import "fmt"

type Student struct {
	name string
	score int
}

func getStudent(studentName string, studentScore int) (student *Student) {
	student = &Student{}
	student.name = studentName
	student.score = studentScore

	return
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(name: %s, score: %d)", s.name, s.score)
}