package trie

import "github.com/donng/Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/BSTMap"

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

func New() *Trie {
	return &Trie{root: generateNode()}
}

// 获得Trie中存储的单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// 向Trie中添加一个新的单词word
func (t *Trie) Add(word string) {
	cur := t.root

	for _, w := range []rune(word) {
		c := string(w)

		if !cur.next.Contains(c) {
			cur.next.Add(c, generateNode())
		}
		cur = cur.next.Get(c).(*Node)
	}

	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}

// 查询单词word是否在Trie中
func (t *Trie) Contains(word string) bool {
	cur := t.root

	for _, w := range []rune(word) {
		c := string(w)
		if !cur.next.Contains(c) {
			return false
		}
		cur = cur.next.Get(c).(*Node)
	}

	return cur.isWord
}
