package main

import "fmt"

type node struct {
	value int
	next map[string]*node
}

type MapSum struct {
	root *node
}

func getNode() *node {
	return &node{next: make(map[string]*node)}
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	n := &node{next: make(map[string]*node)}
	return MapSum{n}
}

func (this *MapSum) Insert(key string, val int)  {
	cur := this.root

	for _, w := range []rune(key) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = getNode()
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

func (this *MapSum) sum(n *node) int {
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
