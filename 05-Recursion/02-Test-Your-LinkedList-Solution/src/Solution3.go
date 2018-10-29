package main

// 虚拟头结点
func RemoveElements3(head *ListNode, val int) *ListNode {
	dummyHead := &ListNodRe{}
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
