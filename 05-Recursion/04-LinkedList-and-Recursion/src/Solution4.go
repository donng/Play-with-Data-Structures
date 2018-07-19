package main

func removeElements4(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements4(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}