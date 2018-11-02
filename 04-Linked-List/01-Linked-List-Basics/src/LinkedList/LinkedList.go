package LinkedList

import "fmt"

type node struct {
	e    interface{}
	next *node
}

func (n *node) String() string {
	return fmt.Sprint(n.e)
}
