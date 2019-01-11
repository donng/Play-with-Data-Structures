package Trie

import "Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/BSTMap"

type Node struct {
	isWord bool
	next   *BSTMap.BSTMap
}

type Trie struct {
	root *Node
	size int
}

func generateNode() *Node {
	return &Node{next: BSTMap.Constructor()}
}

func Constructor() *Trie {
	return &Trie{root: generateNode()}
}

// 获得Trie中存储的单词数量
func (this *Trie) GetSize() int {
	return this.size
}

// 向Trie中添加一个新的单词word
func (this *Trie) Add(word string) {
	cur := this.root

	for _, w := range []rune(word) {
		c := string(w)

		if !cur.next.Contains(c) {
			cur.next.Add(c, generateNode())
		}
		cur = cur.next.Get(c).(*Node)
	}

	if cur.isWord == false {
		cur.isWord = true
		this.size++
	}
}

// 查询单词word是否在Trie中
func (this *Trie) Contains(word string) bool {
	cur := this.root

	for _, w := range []rune(word) {
		c := string(w)
		if !cur.next.Contains(c) {
			return false
		}
		cur = cur.next.Get(c).(*Node)
	}

	return cur.isWord
}
