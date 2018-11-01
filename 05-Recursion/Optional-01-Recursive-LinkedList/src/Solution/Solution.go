package main

import (
	"Play-with-Data-Structures/05-Recursion/Optional-01-Recursive-LinkedList/src/LinkedListR"
	"fmt"
)

type E interface{}

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(e E)
	Pop() E
	Peek() E
}

type LinkedListStack struct {
	listR *LinkedListR.LinkedListR
}

func GetLinkedListStack() *LinkedListStack {
	list := &LinkedListStack{}
	list.listR = LinkedListR.GetLinkedListR()

	return list
}

func (l *LinkedListStack) GetSize() int {
	return l.listR.GetSize()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.listR.IsEmpty()
}

func (l *LinkedListStack) Push(e E) {
	l.listR.AddFirst(e)
}

func (l *LinkedListStack) Pop() E {
	return l.listR.RemoveFirst()
}

func (l *LinkedListStack) Peek() E {
	return l.listR.GetFirst()
}

func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := GetLinkedListStack()

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack.Push(char)
		} else if !stack.IsEmpty() && brackets[char] == stack.Peek() {
			stack.Pop()
		} else {
			return false
		}
	}

	return stack.IsEmpty()
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
}
