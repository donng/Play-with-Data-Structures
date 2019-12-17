package solution

import "github.com/donng/Play-with-Data-Structures/05-Recursion/02-Test-Your-LinkedList-Solution/listnode"

func RemoveElements(head *listnode.ListNode, val int) *listnode.ListNode {
	for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil
	}
	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			cur := prev.Next
			prev.Next = cur.Next
			cur.Next = nil
		} else {
			prev = prev.Next
		}
	}

	return head
}
