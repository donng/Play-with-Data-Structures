package main

import "fmt"

func main() {
	linkedList := getLinkedList()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)
}
