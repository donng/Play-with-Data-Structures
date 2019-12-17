package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/05-Recursion/Optional-01-Recursive-LinkedList/linkedlist"
	"github.com/donng/Play-with-Data-Structures/05-Recursion/Optional-01-Recursive-LinkedList/linkedlistR"
)

func main() {
	// 普通链表测试
	linkedList := linkedlist.New()
	for i := 0; i < 10; i++ {
		linkedList.AddLast(i)
	}
	linkedList.RemoveElement(0)
	linkedList.RemoveElement(6)
	linkedList.RemoveElement(9)
	fmt.Println(linkedList)

	// 递归实现的链表测试
	linkedListR := linkedlistR.New()
	for i := 0; i < 10; i++ {
		linkedListR.AddFirst(i)
	}

	fmt.Println(linkedListR)
	for !linkedListR.IsEmpty() {
		fmt.Println("removed ", linkedListR.RemoveLast())
	}
	fmt.Println(linkedListR)
}
