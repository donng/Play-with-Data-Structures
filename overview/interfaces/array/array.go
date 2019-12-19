package array

type Array interface {
	Add(int, interface{})
	AddFirst(interface{})
	AddLast(interface{})
	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	Find(interface{}) int
	FindAll(interface{}) []int
	Contains(interface{}) bool
	Get(int) interface{}
	Set(int, interface{})
	GetCapacity() int // 获得数组容量
	GetSize() int     // 获得元素个数
	IsEmpty() bool
	RemoveElement(interface{}) bool
}
