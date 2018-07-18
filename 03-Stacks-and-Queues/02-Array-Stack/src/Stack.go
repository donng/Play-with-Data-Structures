package main

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}
