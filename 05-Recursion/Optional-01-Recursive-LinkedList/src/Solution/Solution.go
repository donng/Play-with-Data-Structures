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
	list.listR = LinkedListR.Constructor()

	return list
}

func (ls *LinkedListStack) GetSize() int {
	return ls.listR.GetSize()
}

func (ls *LinkedListStack) IsEmpty() bool {
	return ls.listR.IsEmpty()
}

func (ls *LinkedListStack) Push(e E) {
	ls.listR.AddFirst(e)
}

func (ls *LinkedListStack) Pop() E {
	return ls.listR.RemoveFirst()
}

func (ls *LinkedListStack) Peek() E {
	return ls.listR.GetFirst()
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
