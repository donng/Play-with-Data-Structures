package ArrayStack

import (
	"bytes"
	"fmt"
	"Play-with-Data-Structures/06-Binary-Search-Tree/09-Non-Recursion-Preorder-Traverse-in-BST/src/Array"
)

type ArrayStack struct {
	array *Array.Array
}

func GetArrayStack(capacity int) (arrayStack *ArrayStack) {
	arrayStack = &ArrayStack{}
	arrayStack.array = Array.GetArray(capacity)
	return
}

func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayStack) Push(element interface{}) {
	a.array.AddLast(element)
}

func (a *ArrayStack) Pop() interface{} {
	return a.array.RemoveLast()
}

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
