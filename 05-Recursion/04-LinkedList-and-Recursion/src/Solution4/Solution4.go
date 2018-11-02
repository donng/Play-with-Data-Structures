package Solution4

import "Play-with-Data-Structures/05-Recursion/04-LinkedList-and-Recursion/src/ListNode"

func RemoveElements4(head *ListNode.ListNode, val int) *ListNode.ListNode {
	if head == nil {
		return nil
	}
	head.Next = RemoveElements4(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
