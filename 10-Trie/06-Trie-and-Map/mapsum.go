package main

import "fmt"

type Node struct {
	value int
	next  map[string]*Node
}

type MapSum struct {
	root *Node
}

func generateNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

/** Initialize your data structure here. */
func New() *MapSum {
	return &MapSum{generateNode()}
}

func (s *MapSum) Insert(key string, val int) {
	cur := s.root

	for _, w := range []rune(key) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = generateNode()
		}
		cur = cur.next[c]
	}

	cur.value = val
}

func (s *MapSum) Sum(prefix string) int {
	cur := s.root
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return 0
		}
		cur = cur.next[c]
	}

	return sum(cur)
}

func sum(n *Node) int {
	res := n.value
	for str := range n.next {
		res += sum(n.next[str])
	}
	return res
}

func main() {
	obj := New()
	obj.Insert("apple", 3)
	fmt.Println(obj.Sum("ap"))

	obj.Insert("app", 2)
	fmt.Println(obj.Sum("ap"))
}
