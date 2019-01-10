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

func (this *LinkedListStack) GetSize() int {
	return this.list.GetSize()
}

func (this *LinkedListStack) IsEmpty() bool {
	return this.list.IsEmpty()
}

func (this *LinkedListStack) Push(e interface{}) {
	this.list.AddFirst(e)
}

func (this *LinkedListStack) Pop() interface{} {
	return this.list.RemoveFirst()
}

func (this *LinkedListStack) Peek() interface{} {
	return this.list.GetFirst()
}

func (this *LinkedListStack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Stack: top ")
	buffer.WriteString(this.list.String())

	return buffer.String()
}
