package main

import "fmt"

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
}
