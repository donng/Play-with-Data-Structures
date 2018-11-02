package Stack

type stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}
