package Solution3

import "Play-with-Data-Structures/05-Recursion/04-LinkedList-and-Recursion/src/ListNode"

// 虚拟头结点
func RemoveElements3(head *ListNode.ListNode, val int) *ListNode.ListNode {
	dummyHead := &ListNode.ListNode{}
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
