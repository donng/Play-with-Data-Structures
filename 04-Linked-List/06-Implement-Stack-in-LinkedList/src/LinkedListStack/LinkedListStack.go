package LinkedListStack

import (
	"Play-with-Data-Structures/04-Linked-List/06-Implement-Stack-in-LinkedList/src/LinkedList"
	"bytes"
)

type LinkedListStack struct {
	list *LinkedList.LinkedList
}

func GetLinkedListStack() *LinkedListStack {
	stack := &LinkedListStack{}
	stack.list = LinkedList.GetLinkedList()

	return stack
}

func (l *LinkedListStack) GetSize() int {
	return l.list.GetSize()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedListStack) Push(e interface{}) {
	l.list.AddFirst(e)
}

func (l *LinkedListStack) Pop() interface{} {
	return l.list.RemoveFirst()
}

func (l *LinkedListStack) Peek() interface{} {
	return l.list.GetFirst()
}

func (l *LinkedListStack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Stack: top ")
	buffer.WriteString(l.list.String())

	return buffer.String()
}
