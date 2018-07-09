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
	arrayStack.array = getArray(capacity)
	return
}

func (a *ArrayStack) getSize() int {
	return a.array.getSize()
}

func (a *ArrayStack) isEmpty() bool {
	return a.array.isEmpty()
}

func (a *ArrayStack) push(element interface{}) {
	a.array.addLast(element)
}

func (a *ArrayStack) pop() interface{} {
	return a.array.removeLast()
}

func (a *ArrayStack) peek() interface{} {
	return a.array.getLast()
}

func (a *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < a.array.getSize(); i++ {
		buffer.WriteString(fmt.Sprint(a.array.data[i]))
		if i != a.array.getSize() - 1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}


