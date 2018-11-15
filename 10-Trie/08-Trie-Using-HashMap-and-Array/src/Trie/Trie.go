package Trie

import "Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/BSTMap"

type node struct {
	isWord bool
	next   *BSTMap.BSTMap
}

type Trie struct {
	root *node
	size int
}

func getNode() *node {
	return &node{next: BSTMap.Constructor()}
}

func Constructor() *Trie {
	return &Trie{root: getNode()}
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
			cur.next.Add(c, getNode())
		}
		cur = cur.next.Get(c).(*node)
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
		cur = cur.next.Get(c).(*node)
	}

	return cur.isWord
}

