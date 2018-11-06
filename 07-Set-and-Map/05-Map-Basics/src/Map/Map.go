package Map

type Map interface {
	Add(interface{}, interface{})
	Remove(interface{})
	Get(interface{})
	Set(interface{}, interface{})
	GetSize() int
	IsEmpty() bool
}