package main

func removeElements2(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return head
}
