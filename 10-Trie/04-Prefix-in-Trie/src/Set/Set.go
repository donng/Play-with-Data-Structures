package Set

type E interface{}

type Set interface {
	Add(E)
	Remove(E)
	Contains(E) bool
	GetSize() int
	IsEmpty() bool
}
