package arraystack

import (
	"bytes"
	"fmt"
	"github.com/donng/Play-with-Data-Structures/03-Stacks-and-Queues/04-More-about-Leetcode/array"
)

type ArrayStack struct {
	array *array.Array
}

func Constructor(capacity int) *ArrayStack {
	return &ArrayStack{
		array: array.New(capacity),
	}
}

func (as *ArrayStack) GetSize() int {
	return as.array.GetSize()
}

func (as *ArrayStack) IsEmpty() bool {
	return as.array.IsEmpty()
}

// 压入栈顶元素
func (as *ArrayStack) Push(element interface{}) {
	as.array.AddLast(element)
}

// 弹出栈顶元素
func (as *ArrayStack) Pop() interface{} {
	return as.array.RemoveLast()
}

// 查看栈顶元素
func (as *ArrayStack) Peek() interface{} {
	return as.array.GetLast()
}

func (as *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < as.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(as.array.Get(i)))
		if i != as.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
