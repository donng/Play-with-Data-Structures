package Student

import (
	"hash/fnv"
	"strconv"
)

type Student struct {
	grade     int
	cls       int
	firstName string
	lastName  string
}

func Constructor(grade int, cls int, firstName string, lastName string) *Student {
	return &Student{grade, cls, firstName, lastName}
}

func (s *Student) HashCode() int {
	B := 31
	hash := 0
	hash = hash*B + int(HashCode(s.grade))
	hash = hash*B + int(HashCode(s.cls))
	hash = hash*B + int(HashCode(s.firstName))
	hash = hash*B + int(HashCode(s.lastName))

	return hash
}

func HashCode(source interface{}) uint64 {
	switch source.(type) {
	case int:
		source = strconv.Itoa(source.(int))
	case float64:
		source = strconv.FormatFloat(source.(float64), 'f', 6, 64)
	}

	h := fnv.New64a()
	h.Write([]byte(source.(string)))
	return h.Sum64()
}

func (s *Student) equals(another Student) bool {
	return s.grade == another.grade && s.cls == another.cls &&
		s.firstName == another.firstName && s.lastName == another.lastName
}
