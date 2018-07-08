package main

import (
	"fmt"
)

func main() {
	arr := getArray(20)
	for i := 0; i < 10; i++ {
		arr.addLast(i)
	}
	fmt.Println(arr)

	arr.add(1, 100)
	fmt.Println(arr)

	arr.addFirst(-1)
	fmt.Println(arr)
	// [-1, 0, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9]

	arr.remove(2)
	fmt.Println(arr)

	arr.removeElement(4)
	fmt.Println(arr)

	arr.removeFirst()
	fmt.Println(arr)

	arr.addLast(1)
	indexes := arr.findAll(1)
	fmt.Println(arr, indexes)

	isRemove := arr.removeAllElement(1)
	fmt.Println(arr, isRemove)

	// 测试 student 类型
	students := getArray(10)
	students.addLast(Student{"Alice", 100})
	students.addLast(Student{"Bob", 66})
	students.addLast(Student{"Charlie", 88})
	fmt.Println(students)
}
