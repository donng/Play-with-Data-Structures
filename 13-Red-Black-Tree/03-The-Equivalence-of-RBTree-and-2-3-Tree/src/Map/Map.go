package Map

type Map interface {
	Add(interface{}, interface{})
	Remove(interface{}) interface{}
	Contains(interface{}) bool
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	GetSize() int
	IsEmpty() bool
}
