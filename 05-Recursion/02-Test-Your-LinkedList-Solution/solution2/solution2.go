package solution2

import "github.com/donng/Play-with-Data-Structures/05-Recursion/02-Test-Your-LinkedList-Solution/listnode"

func RemoveElements(head *listnode.ListNode, val int) *listnode.ListNode {
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
