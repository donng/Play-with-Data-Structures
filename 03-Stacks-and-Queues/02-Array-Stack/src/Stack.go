package main

type stack interface {
	getSize() int
	isEmpty() bool
	push(interface{})
	pop() interface{}
	peek() interface{}
}
