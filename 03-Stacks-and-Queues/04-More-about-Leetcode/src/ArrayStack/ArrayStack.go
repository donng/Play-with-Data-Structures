package ArrayStack

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/04-More-about-Leetcode/src/Array"
	"bytes"
	"fmt"
)

type ArrayStack struct {
	array *Array.Array
}

func GetArrayStack(capacity int) (arrayStack *ArrayStack) {
	arrayStack = &ArrayStack{}
	arrayStack.array = Array.GetArray(capacity)
	return
}

func (s *ArrayStack) GetSize() int {
	return s.array.GetSize()
}

func (s *ArrayStack) IsEmpty() bool {
	return s.array.IsEmpty()
}

func (s *ArrayStack) Push(element interface{}) {
	s.array.AddLast(element)
}

func (s *ArrayStack) Pop() interface{} {
	return s.array.RemoveLast()
}

func (s *ArrayStack) Peek() interface{} {
	return s.array.GetLast()
}

func (s *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < s.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(s.array.Get(i)))
		if i != s.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
