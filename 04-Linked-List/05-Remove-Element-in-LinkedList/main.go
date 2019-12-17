package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/04-Linked-List/05-Remove-Element-in-LinkedList/linkedlist"
)

func main() {
	linkedList := linkedlist.New()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)

	linkedList.Remove(2)
	fmt.Println(linkedList)

	linkedList.RemoveFirst()
	fmt.Println(linkedList)

	linkedList.RemoveLast()
	fmt.Println(linkedList)
}
