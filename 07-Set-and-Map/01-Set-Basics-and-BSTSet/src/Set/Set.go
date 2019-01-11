package Set

type Set interface {
	Add(interface{})
	Remove(interface{})
	Contains(interface{}) bool
	GetSize() int
	IsEmpty() bool
}
