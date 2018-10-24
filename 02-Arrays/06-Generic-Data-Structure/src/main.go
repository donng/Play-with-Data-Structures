package main

import (
	"Play-with-Data-Structures/02-Arrays/06-Generic-Data-Structure/src/Array"
	"Play-with-Data-Structures/02-Arrays/06-Generic-Data-Structure/src/Student"
	"fmt"
)

func main() {
	arr := Array.GetArray(20)

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
	students := Array.GetArray(10)
	students.AddLast(Student.GetStudent("Alice", 100))
	students.AddLast(Student.GetStudent("Bob", 66))
	students.AddLast(Student.GetStudent("Charlie", 88))
	fmt.Println(students)
}
