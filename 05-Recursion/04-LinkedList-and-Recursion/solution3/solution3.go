package solution3

import "github.com/donng/Play-with-Data-Structures/05-Recursion/04-LinkedList-and-Recursion/listnode"

// 虚拟头结点
func RemoveElements(head *listnode.ListNode, val int) *listnode.ListNode {
	dummyHead := &listnode.ListNode{}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return dummyHead.Next
}
