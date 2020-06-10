package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/02-Arrays/06-Generic-Data-Structures/array"
	"github.com/donng/Play-with-Data-Structures/02-Arrays/06-Generic-Data-Structures/student"
)

func main() {
	arr := array.New(20)

	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)
	// [-1, 0, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9]

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	arr.AddLast(1)
	indexes := arr.FindAll(1)
	fmt.Println(arr, indexes)

	isRemove := arr.RemoveAllElement(1)
	fmt.Println(arr, isRemove)

	// 测试 student 类型
	students := array.New(10)
	students.AddLast(student.New("Alice", 100))
	students.AddLast(student.New("Bob", 66))
	students.AddLast(student.New("Charlie", 88))
	fmt.Println(students)
}
