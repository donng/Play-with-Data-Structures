package Student

import "fmt"

type Student struct {
	name string
	score int
}

func GetStudent(name string, score int) (student *Student) {
	student = &Student{}
	student.name = name
	student.score = score

	return
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(name: %s, score: %d)", s.name, s.score)
}