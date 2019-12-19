package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/14-Hash-Table/03-Hash-Function-in-Go/student"
)

func main() {
	a := 42
	fmt.Println(student.HashCode(a))

	b := -42
	fmt.Println(student.HashCode(b))

	c := 3.1415926
	fmt.Println(student.HashCode(c))

	d := "imooc"
	fmt.Println(student.HashCode(d))

	//fmt.Println(math.MaxInt64 + 1)
	s := student.New(3, 2, "Bobo", "Liu")
	fmt.Println(s.HashCode())

	scores := make(map[*student.Student]int)
	scores[s] = 100

	student2 := student.New(3, 2, "Bobo", "Liu")
	fmt.Println(student2.HashCode())
}
