package ArrayStack

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/04-More-about-Leetcode/src/Array"
	"bytes"
	"fmt"
)

type ArrayStack struct {
	array *Array.Array
}

func Constructor(capacity int) *ArrayStack {
	return &ArrayStack{
		array: Array.Constructor(capacity),
	}
}

func (this *ArrayStack) GetSize() int {
	return this.array.GetSize()
}

func (this *ArrayStack) IsEmpty() bool {
	return this.array.IsEmpty()
}

// 压入栈顶元素
func (this *ArrayStack) Push(element interface{}) {
	this.array.AddLast(element)
}

// 弹出栈顶元素
func (this *ArrayStack) Pop() interface{} {
	return this.array.RemoveLast()
}

// 查看栈顶元素
func (this *ArrayStack) Peek() interface{} {
	return this.array.GetLast()
}

func (this *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < this.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(this.array.Get(i)))
		if i != this.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
