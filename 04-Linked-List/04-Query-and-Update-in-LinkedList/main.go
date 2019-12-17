package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/04-Linked-List/04-Query-and-Update-in-LinkedList/linkedlist"
)

func main() {
	linkedList := linkedlist.New()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)
}
