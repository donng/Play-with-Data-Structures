package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/05-Recursion/Optional-01-Recursive-LinkedList/linkedlistR"
)

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
}

type LinkedListStack struct {
	listR *linkedlistR.LinkedListR
}

func GetLinkedListStack() *LinkedListStack {
	list := &LinkedListStack{}
	list.listR = linkedlistR.New()

	return list
}

func (ls *LinkedListStack) GetSize() int {
	return ls.listR.GetSize()
}

func (ls *LinkedListStack) IsEmpty() bool {
	return ls.listR.IsEmpty()
}

func (ls *LinkedListStack) Push(e interface{}) {
	ls.listR.AddFirst(e)
}

func (ls *LinkedListStack) Pop() interface{} {
	return ls.listR.RemoveFirst()
}

func (ls *LinkedListStack) Peek() interface{} {
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
