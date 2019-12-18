package merger

type Merger interface {
	merge(interface{}, interface{}) interface{}
}
