package main

type Queue interface {
	getSize() int
	isEmpty() bool
	enqueue(interface{})
	dequeue() interface{}
	getFront() interface{}
}