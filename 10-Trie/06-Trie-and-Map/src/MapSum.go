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

func (this *MapSum) Insert(key string, val int) {
	cur := this.root

	for _, w := range []rune(key) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = generateNode()
		}
		cur = cur.next[c]
	}

	cur.value = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return 0
		}
		cur = cur.next[c]
	}

	return this.sum(cur)
}

func (this *MapSum) sum(n *Node) int {
	res := n.value
	for s := range n.next {
		res += this.sum(n.next[s])
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
