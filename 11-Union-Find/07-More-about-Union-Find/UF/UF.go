package UF

type UF interface {
	GetSize() int
	IsConnected(p int, q int) bool
	UnionElements(p int, q int)
}
