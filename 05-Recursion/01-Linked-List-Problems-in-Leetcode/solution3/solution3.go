package solution3

import (
	"github.com/donng/Play-with-Data-Structures/05-Recursion/01-Linked-List-Problems-in-Leetcode/listnode"
)

/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/
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
