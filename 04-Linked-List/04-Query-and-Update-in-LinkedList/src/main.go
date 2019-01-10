package main

import (
	"Play-with-Data-Structures/04-Linked-List/04-Query-and-Update-in-LinkedList/src/LinkedList"
	"fmt"
)

func main() {
	linkedList := LinkedList.Constructor()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)
}
