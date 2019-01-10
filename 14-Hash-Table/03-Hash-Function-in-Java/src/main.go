package main

import (
	"Play-with-Data-Structures/14-Hash-Table/03-Hash-Function-in-Java/src/Student"
	"fmt"
)

func main() {
	a := 42
	fmt.Println(Student.HashCode(a))

	b := -42
	fmt.Println(Student.HashCode(b))

	c := 3.1415926
	fmt.Println(Student.HashCode(c))

	d := "imooc"
	fmt.Println(Student.HashCode(d))

	//fmt.Println(math.MaxInt64 + 1)
	student := Student.Constructor(3, 2, "Bobo", "Liu")
	fmt.Println(student.HashCode())

	scores := make(map[*Student.Student]int)
	scores[student] = 100

	student2 := Student.Constructor(3, 2, "Bobo", "Liu")
	fmt.Println(student2.HashCode())
}
