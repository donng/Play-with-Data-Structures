package main

import (
	"bytes"
	"fmt"
)

type ArrayStack struct {
	array *Array
}

func getArrayStack(capacity int) (arrayStack *ArrayStack) {
	arrayStack = &ArrayStack{}
	arrayStack.array = GetArray(capacity)
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
		buffer.WriteString(fmt.Sprint(a.array.data[i]))
		if i != a.array.GetSize() - 1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}


