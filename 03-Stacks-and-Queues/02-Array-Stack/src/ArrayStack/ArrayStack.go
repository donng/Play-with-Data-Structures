package ArrayStack

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/02-Array-Stack/src/Array"
	"bytes"
	"fmt"
)

type ArrayStack struct {
	array *Array.Array
}

func GetArrayStack(capacity int) *ArrayStack {
	arrayStack := &ArrayStack{}
	arrayStack.array = Array.GetArray(capacity)
	return arrayStack
}

func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

// 压入栈顶元素
func (a *ArrayStack) Push(element interface{}) {
	a.array.AddLast(element)
}

// 弹出栈顶元素
func (a *ArrayStack) Pop() interface{} {
	return a.array.RemoveLast()
}

// 查看栈顶元素
func (a *ArrayStack) Peek() interface{} {
	return a.array.GetLast()
}

func (a *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < a.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(a.array.Get(i)))
		if i != a.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
