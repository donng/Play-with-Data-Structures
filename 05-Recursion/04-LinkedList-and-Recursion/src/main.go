package main

import (
	"fmt"
	"Play-with-Data-Structures/05-Recursion/04-LinkedList-and-Recursion/src/ListNode"
	"Play-with-Data-Structures/05-Recursion/04-LinkedList-and-Recursion/src/Solution4"
)

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := ListNode.GetListNode(nums)
	fmt.Println(head)

	res := Solution4.RemoveElements4(head, 6)
	fmt.Println(res)
}
