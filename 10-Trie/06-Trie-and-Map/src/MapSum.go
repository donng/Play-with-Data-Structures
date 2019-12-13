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
func Constructor() MapSum {
	return MapSum{
		root: generateNode(),
	}
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

	return s.sum(cur)
}

func (s *MapSum) sum(n *Node) int {
	res := n.value
	for s := range n.next {
		res += s.sum(n.next[s])
	}
	return res
}

func main() {
	obj := Constructor()
	obj.Insert("apple", 3)
	fmt.Println(obj.Sum("ap"))

	obj.Insert("app", 2)
	fmt.Println(obj.Sum("ap"))
}
