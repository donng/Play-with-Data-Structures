package solution4

import "github.com/donng/Play-with-Data-Structures/05-Recursion/04-LinkedList-and-Recursion/listnode"

func RemoveElements(head *listnode.ListNode, val int) *listnode.ListNode {
	if head == nil {
		return nil
	}
	head.Next = RemoveElements(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
