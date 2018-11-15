package Trie3

type node struct {
	isWord bool
	next   [26]*node
}

type Trie struct {
	root *node
	size int
}

func getNode() *node {
	return &node{next: [26]*node{}}
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

	for _, w := range word {
		offset := w - 'a'
		if offset > 26 {
			offset -= 26
		}

		if cur.next[offset] == nil {
			cur.next[offset] = getNode()
		}
		cur = cur.next[offset]
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
		offset := w - 'a'
		if cur.next[offset] == nil {
			return false
		}
		cur = cur.next[offset]
	}

	return cur.isWord
}
