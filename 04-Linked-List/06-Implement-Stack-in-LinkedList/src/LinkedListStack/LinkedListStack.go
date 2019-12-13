package LinkedListStack

import (
	"Play-with-Data-Structures/04-Linked-List/06-Implement-Stack-in-LinkedList/src/LinkedList"
	"bytes"
)

type LinkedListStack struct {
	list *LinkedList.LinkedList
}

func Constructor() *LinkedListStack {
	return &LinkedListStack{
		LinkedList.Constructor(),
	}
}

func (ls *LinkedListStack) GetSize() int {
	return ls.list.GetSize()
}

func (ls *LinkedListStack) IsEmpty() bool {
	return ls.list.IsEmpty()
}

func (ls *LinkedListStack) Push(e interface{}) {
	ls.list.AddFirst(e)
}

func (ls *LinkedListStack) Pop() interface{} {
	return ls.list.RemoveFirst()
}

func (ls *LinkedListStack) Peek() interface{} {
	return ls.list.GetFirst()
}

func (ls *LinkedListStack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Stack: top ")
	buffer.WriteString(ls.list.String())

	return buffer.String()
}
