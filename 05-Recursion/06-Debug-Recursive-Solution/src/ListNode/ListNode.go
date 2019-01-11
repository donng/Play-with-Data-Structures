package ListNode

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表节点的构造函数
// 使用 arr 为参数，创建一个链表，当前的 ListNode 为链表头结点
func GetListNode(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		panic("arr can not be empty")
	}

	head := &ListNode{}
	head.Val = arr[0]

	prev := head
	for i := 1; i < len(arr); i++ {
		prev.Next = &ListNode{Val: arr[i]}
		prev = prev.Next
	}

	return head
}

func (l *ListNode) String() string {
	var buffer bytes.Buffer

	cur := l
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", cur.Val))
		cur = cur.Next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
