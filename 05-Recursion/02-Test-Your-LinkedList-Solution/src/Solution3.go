package main

// 虚拟头结点
func removeElements3(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
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
